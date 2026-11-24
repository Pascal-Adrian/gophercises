package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"testing/synctest"
	"time"
)

func TestReadToChan(t *testing.T) {
	var table = []struct {
		name string
		in   io.Reader
		out  string
	}{
		{"string with new line", strings.NewReader("input\n"), "input\n"},
		{"empty string", strings.NewReader(""), ""},
		{"no trailing new line", strings.NewReader("10"), "10"},
		{"error reader", iotest.ErrReader(errors.New("error")), ""},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			channel := make(chan string, 1)

			reader := bufio.NewReader(tc.in)

			go readToChan(channel, reader)

			value, _ := <-channel

			if value != tc.out {
				t.Errorf("got %q; want %q", value, tc.out)
			}
		})
	}
}

func TestCheckAnswer(t *testing.T) {
	var table = []struct {
		name string
		in   [2]string
		out  bool
	}{
		{"simple strings", [2]string{"input", "input"}, true},
		{"empty strings", [2]string{"", ""}, true},
		{"first string with uppercase", [2]string{"InPuT", "input"}, true},
		{"first string with trailing and leading spaces", [2]string{"     input     ", "input"}, true},
		{"first string with trailing and leading spaces, tabs, new lines", [2]string{"     \t\t\t\ninput\n\n\n\n\n     ", "input"}, true},
		{"second string with uppercase", [2]string{"input", "InPuT"}, true},
		{"second string with trailing and leading spaces", [2]string{"input", "     input     "}, true},
		{"second string with trailing and leading spaces, tabs, new lines", [2]string{"input", "     \t\t\t\ninput\n\n\n\n\n     "}, true},
		{"different strings", [2]string{"input", "inaput"}, false},
		{"partial matching strings", [2]string{"input", "inpu"}, false},
		{"first string with space", [2]string{"inp ut", "input"}, false},
		{"second string with space", [2]string{"input", "inp ut"}, false},
		{"both string with same space", [2]string{"inp ut", "inp ut"}, true},
		{"both string with different space", [2]string{"in put", "inp ut"}, false},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			result := checkAnswer(tc.in[0], tc.in[1])

			if result != tc.out {
				t.Errorf("checkAnswer(%q, %q) = %t; want %t", tc.in[0], tc.in[1], result, tc.out)
			}
		})
	}
}

var testProblems = [][]string{{"1+1", "2"}, {"2+2", "4"}, {"3+3", "6"}}

func TestAskCount(t *testing.T) {
	var table = []struct {
		name string
		in   string
		out  int
	}{
		{"all correct", "2\n4\n6\n", 3},
		{"all wrong", "1\n1\n1\n", 0},
		{"mixed", "2\n5\n6\n", 2},
		{"answers matched to the right question", "4\n6\n2\n", 0},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tc.in))

			correct, err := ask(reader, testProblems, time.Second)

			if err != nil {
				t.Errorf("ask(%q, %v) = %d, %v; want %d, nil", tc.in, testProblems, correct, err, tc.out)
			}

			if correct != tc.out {
				t.Errorf("ask(%q, %v) = %d; want %d", tc.in, testProblems, correct, tc.out)
			}
		})
	}
}

func TestAskTimeout(t *testing.T) {
	type answer struct {
		at   time.Duration
		text string
	}

	const timeout = 10 * time.Second

	var table = []struct {
		name        string
		in          []answer
		wantElapsed time.Duration
		wantCorrect int
		wantErr     bool
	}{
		{"no answer", nil, 10 * time.Second, 0, true},
		{"2 answers then silence", []answer{{1 * time.Second, "2\n"}, {2 * time.Second, "4\n"}}, 10 * time.Second, 2, true},
		{"wrong answers then silence", []answer{{1 * time.Second, "3\n"}}, 10 * time.Second, 0, true},
		{"total time, not per question", []answer{{4 * time.Second, "2\n"}, {8 * time.Second, "4\n"}, {12 * time.Second, "6\n"}}, 10 * time.Second, 2, true},
		{"last just before deadline", []answer{{4 * time.Second, "2\n"}, {8 * time.Second, "4\n"}, {timeout - time.Nanosecond, "6\n"}}, 10*time.Second - time.Nanosecond, 3, false},
		{"last just after deadline", []answer{{4 * time.Second, "2\n"}, {8 * time.Second, "4\n"}, {timeout + time.Nanosecond, "6\n"}}, 10 * time.Second, 2, true},
		{"partial line when time runs out", []answer{{4 * time.Second, "2\n"}, {8 * time.Second, "4"}}, 10 * time.Second, 1, true},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				pr, pw := io.Pipe()
				defer pw.Close()

				go func() {
					var now time.Duration

					for _, a := range tc.in {
						time.Sleep(a.at - now)
						now = a.at

						if _, err := io.WriteString(pw, a.text); err != nil {
							return
						}
					}
				}()

				start := time.Now()
				correct, err := ask(bufio.NewReader(pr), testProblems, timeout)
				elapsed := time.Since(start)

				time.Sleep(time.Minute)

				if err != nil && !tc.wantErr {
					t.Errorf("ask(%v, %v, %v) = %d, %v; want %d, nil", tc.in, testProblems, timeout, correct, err, tc.wantCorrect)
				}

				if tc.wantErr && err == nil {
					t.Errorf("ask(%v, %v, %v) = %d, nil; want %v, %v", tc.in, testProblems, timeout, correct, tc.wantCorrect, errTimeout)
				}

				if correct != tc.wantCorrect {
					t.Errorf("ask(%v, %v, %v) = %d; want %d", tc.in, testProblems, timeout, correct, tc.wantCorrect)
				}

				if elapsed != tc.wantElapsed {
					t.Errorf("ask(%v, %v, %v) = %v; want %v", tc.in, testProblems, timeout, elapsed, tc.wantElapsed)
				}
			})
		})
	}
}

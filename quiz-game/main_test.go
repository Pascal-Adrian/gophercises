package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
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

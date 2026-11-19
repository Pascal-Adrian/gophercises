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

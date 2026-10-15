package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadRESP(t *testing.T) {
	cases := []struct {
		input string
		want  any
	}{
		{"+OK\r\n", "OK"}, {":17\r\n", int64(17)}, {"$3\r\nhey\r\n", "hey"}, {"$-1\r\n", nil},
	}
	for _, tc := range cases {
		got, err := readRESP(bufio.NewReader(strings.NewReader(tc.input)))
		if err != nil || got != tc.want {
			t.Fatalf("input=%q got=%v err=%v", tc.input, got, err)
		}
	}
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"command", "one", "two", "three"}, "one two three\none two three\n"},
		{[]string{"command", "a", "b", "c"}, "a b c\na b c\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		out = new(bytes.Buffer)
		main()

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("main() = %q, want %q", got, test.want)
		}
	}
}

func BenchmarkEcho(b *testing.B) {
	args := make([]string, 1000)
	for i := 0; i < len(args); i++ {
		args[i] = "hoge"
	}

	for i := 0; i < b.N; i++ {
		echo(args)
	}
}

func BenchmarkEchoEfficiency(b *testing.B) {
	args := make([]string, 1000)
	for i := 0; i < len(args); i++ {
		args[i] = "hoge"
	}

	for i := 0; i < b.N; i++ {
		echoEfficiency(args)
	}
}

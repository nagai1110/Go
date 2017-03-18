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
		{[]string{"command"}, ""},
		{[]string{"command", "one", "two", "three"}, "1 one\n2 two\n3 three\n"},
		{[]string{"command", "a", "b", "c"}, "1 a\n2 b\n3 c\n"},
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

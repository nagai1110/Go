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
		{[]string{}, "\n"},
		{[]string{"one", "two", "three"}, "one two three\n"},
		{[]string{"a", "b", "c"}, "a b c\n"},
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

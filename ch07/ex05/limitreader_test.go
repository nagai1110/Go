// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var tests = []struct {
		str   string
		limit int64
		buf   int
		want  []byte
	}{
		{"123456789", 1, 0, []byte("")},
		{"123456789", 3, 1, []byte("123")},
		{"123456789", 10, 2, []byte("123456789")},
		{"123456789", 10, 10, []byte("123456789")},
	}

	for _, test := range tests {
		lr := LimitReader(strings.NewReader(test.str), test.limit)
		buf := make([]byte, test.buf)

		lr.Read(buf)
		for i := 0; i < len(test.want) && i < len(buf); i++ {
			if buf[i] != test.want[i] {
				t.Errorf("Read() = %v, want %v", buf, test.want)
				break
			}
		}
	}
}

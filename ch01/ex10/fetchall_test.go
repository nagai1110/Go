// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFetchAll(t *testing.T) {
	var tests = []struct {
		args []string
		want []string
	}{
		{[]string{"command", "http://gopl.io"}, []string{"file_1.txt", "file_2.txt"}},
		{[]string{"command", "http://ricoh.co.jp"}, []string{"file_3.txt", "file_4.txt"}},
		{[]string{"command", "http://www.gundam.info"}, []string{"file_5.txt", "file_6.txt"}},
	}

	for _, test := range tests {
		os.Args = test.args
		main()

		for _, filename := range test.want {
			bytes, err := ioutil.ReadFile(filename)
			if err != nil {
				t.Errorf("%s isn't created", filename)
			}

			if len(bytes) <= 0 {
				t.Errorf("%s is empty file", filename)
			}
		}
	}
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // テスト中は変更される

func main() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout // テスト中は変更される

func main() {
	echo(os.Args)
	echoEfficiency(os.Args)
}

func echo(args []string) {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Fprintln(out, s)
}

func echoEfficiency(args []string) {
	fmt.Fprintln(out, strings.Join(os.Args[1:], " "))
}

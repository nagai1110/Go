// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // テスト中は変更される

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Fprintln(out, fmt.Sprintf("%d %s", i, os.Args[i]))
	}
}

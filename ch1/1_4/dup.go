// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout // for Test

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	print(counts)
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

// 文字列順にソートして出力
func print(counts map[string][]string) {
	var sortedKeys []string
	for key := range counts {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		filenames := counts[key]
		if len(filenames) > 1 {
			fmt.Fprintf(out, "%s\t%s\n", key, filenames)
		}
	}
}

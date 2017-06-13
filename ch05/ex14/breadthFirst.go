// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func walkDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var paths []string
	for _, file := range files {
		absPath, _ := filepath.Abs(filepath.Join(dir, file.Name()))
		if file.IsDir() {
			paths = append(paths, absPath)
		} else {
			fmt.Println(absPath)
		}
	}

	return paths
}

func main() {
	breadthFirst(walkDir, os.Args[1:])
}

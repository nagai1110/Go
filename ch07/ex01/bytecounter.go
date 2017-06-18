// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewBuffer(p))

	var count int
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		count++
	}

	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewBuffer(p))

	var count int
	for sc.Scan() {
		count++
	}

	*c += LineCounter(count)
	return count, nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("1 2 3"))
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("1\n2 2 \n3 3 3\n4 4 4 4\n5 5 5 5 5"))
	fmt.Println(lc)
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var runeTypeCounts = make(map[string]int)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		countRuneType(r)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Println("\nrune type\tcount")
	for t, n := range runeTypeCounts {
		fmt.Printf("%s\t%d\n", t, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countRuneType(r rune) {
	if unicode.IsDigit(r) {
		runeTypeCounts["Digit"]++
	}
	if unicode.IsGraphic(r) {
		runeTypeCounts["Graphic"]++
	}
	if unicode.IsLetter(r) {
		runeTypeCounts["Letter"]++
	}
	if unicode.IsMark(r) {
		runeTypeCounts["Mark"]++
	}
	if unicode.IsNumber(r) {
		runeTypeCounts["Number"]++
	}
	if unicode.IsPunct(r) {
		runeTypeCounts["Punct"]++
	}
	if unicode.IsSpace(r) {
		runeTypeCounts["Space"]++
	}
}

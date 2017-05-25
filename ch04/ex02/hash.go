// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var sha = flag.String("sha", "sha256", "hoge")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadSlice('\n')

	switch *sha {
	case "sha512":
		fmt.Printf("%s : %x\n", "sha512", sha512.Sum512([]byte(word)))
	case "sha384":
		fmt.Printf("%s : %x\n", "sha384", sha512.Sum384([]byte(word)))
	case "sha256":
		fallthrough
	default:
		fmt.Printf("%s : %x\n", "sha256", sha256.Sum256([]byte(word)))
	}
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var out io.Writer = os.Stdout // for Test

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Fprintf(out, "status code: %d\n", resp.StatusCode)
	fmt.Fprintf(out, "%s", b)
}

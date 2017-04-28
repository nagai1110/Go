// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if isSucceeded := fetch(url); !isSucceeded {
			os.Exit(1)
		}
	}
}

func fetch(url string) bool {
	resp, err := http.Get(createUrl(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return false
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return false
	}

	fmt.Printf("%s", b)
	return true
}

func createUrl(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}

	return "http://" + url
}

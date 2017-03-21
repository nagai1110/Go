// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var fileIndex int = 0

func main() {
	reportFetch(os.Args[1:])
	reportFetch(os.Args[1:])
}

func reportFetch(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	fileIndex++
	filename := fmt.Sprintf("file_%d.txt", fileIndex)
	out, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(out, resp.Body)
	out.Close()
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

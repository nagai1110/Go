// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type linkInfo struct {
	url   string
	depth int
}

func crawl(link linkInfo) []linkInfo {
	fmt.Println(link.url)
	urls, err := links.Extract(link.url)
	if err != nil {
		log.Print(err)
	}

	return createLinkInfos(urls, link.depth+1)
}

func createLinkInfos(urls []string, depth int) []linkInfo {
	var linkInfos []linkInfo
	for _, url := range urls {
		linkInfos = append(linkInfos, linkInfo{url: url, depth: depth})
	}

	return linkInfos
}

func main() {
	var depth int
	flag.IntVar(&depth, "depth", 1, "depth")
	flag.Parse()

	worklist := make(chan []linkInfo)
	unseenLinks := make(chan linkInfo)

	go func() {
		worklist <- createLinkInfos(os.Args[1:], 0)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] && link.depth <= depth {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

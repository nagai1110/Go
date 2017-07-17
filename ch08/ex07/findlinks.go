// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

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

func getRootDomains(urls []string) []string {
	var domains []string
	for _, u := range urls {
		url, err := url.Parse(u)
		if err == nil {
			domains = append(domains, url.Host)
		}
	}

	return domains
}

func isRootDomain(urlStr string, domains []string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	for _, domain := range domains {
		if u.Host == domain {
			return true
		}
	}

	return false
}

func createFile(urlStr string) {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	dir, file := path.Split(u.Path)
	_, err = os.Stat(dir)
	if err == nil {
		return
	}

	if err := os.MkdirAll("result/"+u.Host+dir, 0777); err != nil {
		fmt.Println(err)
		return
	}

	if len(file) == 0 {
		return
	}

	response, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.OpenFile("result/"+u.Host+dir+file, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.Write(body)
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

	rootDomains := getRootDomains(os.Args[1:])
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] && link.depth <= depth && isRootDomain(link.url, rootDomains) {
				seen[link.url] = true
				createFile(link.url)
				unseenLinks <- link
			}
		}
	}
}

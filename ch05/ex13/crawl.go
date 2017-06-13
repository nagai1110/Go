// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"gopl.io/ch5/links"
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

func crawl(urlStr string) []string {
	fmt.Println(urlStr)
	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}

	u, _ := url.Parse(urlStr)
	for _, link := range list {
		linkURL, _ := url.Parse(link)
		if linkURL.Host == u.Host {
			createFile(u)
		}
	}

	return list
}

func createFile(u *url.URL) {
	dir, file := path.Split(u.Path)
	_, err := os.Stat(dir)
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
	breadthFirst(crawl, os.Args[1:])
}

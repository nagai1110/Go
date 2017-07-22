// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func mirroredQuery(urls []string) {
	resps := make(chan struct{})
	cancel := make(chan struct{})

	cancelled := func() bool {
		select {
		case <-cancel:
			return true
		default:
			return false
		}
	}

	for _, url := range urls {
		go func() {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			req.Cancel = cancel

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			if !cancelled() {
				io.Copy(ioutil.Discard, resp.Body)
				fmt.Printf("%s %d\n", url, resp.StatusCode)
			}

			close(cancel)
			resps <- struct{}{}
		}()
	}

	<-resps
}

func main() {
	mirroredQuery(os.Args[1:])
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"net/http"
	"testing"
	"time"
)

func TestSurfaceServer(t *testing.T) {
	// サーバーにアクセス可能かどうかのみ検査
	var tests = []struct {
		url  string
		want int
	}{
		{"http://localhost:8000", 200},
	}

	// 非同期でmain()を実行してサーバーをたてる
	go main()
	time.Sleep(1 * time.Second)

	for _, test := range tests {
		resp, err := http.Get(test.url)
		if err != nil || resp.StatusCode != 200 {
			t.Errorf("Error. url: %s err: %v\n", test.url, err)
		}
	}
}

// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

var sleeper = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 100秒間スリープさせる
	time.Sleep(100 * time.Second)
})

func TestFetchAll(t *testing.T) {
	// 応答しないWebサイトが分からなかったのでテストサーバーでスリープさせてみる
	testServer := httptest.NewServer(sleeper)
	defer testServer.Close()

	var tests = []struct {
		args []string
		want []string
	}{
		{[]string{
			"command",
			"http://google.com",
			"http://youtube.com",
			"http://facebook.com",
			"http://baidu.com",
			"http://wikipedia.org",
			testServer.URL, // 途中でテストサーバーにアクセスさせてみる
			"http://yahoo.com",
			"http://google.co.in",
			"http://qq.com",
			"http://taobao.com",
			"http://tmall.com"},

			[]string{
				"http://google.com: 200",
				"http://youtube.com: 200",
				"http://facebook.com: 200",
				"http://baidu.com: 200",
				"http://wikipedia.org: 200",
				fmt.Sprintf("%s: 200", testServer.URL),
				"http://yahoo.com: 200",
				"http://google.co.in: 200",
				"http://qq.com: 200",
				"http://taobao.com: 200",
				"http://tmall.com: 200",
			}},
	}

	for _, test := range tests {
		os.Args = test.args
		statusCodeWriter = new(bytes.Buffer)
		main()

		got := statusCodeWriter.(*bytes.Buffer).String()
		for _, result := range test.want {
			if !strings.Contains(got, result) {
				t.Errorf("main() = %q, want = %q", got, test.want)
			}
		}
	}
}

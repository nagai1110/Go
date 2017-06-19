// Copyright © 2017 Takaya.Nagai All Rights Res
package main

import (
	"io"
)

type limitReader struct {
	reader io.Reader
	limit  int64
	count  int
}

func (lr *limitReader) Read(p []byte) (int, error) {
	if len(p) == 0 || int64(lr.count) >= lr.limit {
		return 0, io.EOF
	}

	limit := int(lr.limit - int64(lr.count))
	if limit > len(p) {
		limit = len(p)
	}

	n, err := lr.reader.Read(p[:limit])
	lr.count += limit
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := new(limitReader)
	lr.reader = r
	lr.limit = n
	lr.count = 0
	return lr
}

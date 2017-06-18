// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	writer io.Writer
	count  int64
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	count, err := cw.writer.Write(p)
	cw.count += int64(count)
	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := new(CountWriter)
	cw.writer = w

	return cw, &cw.count
}

func main() {
	w, c := CountingWriter(os.Stdout)
	w.Write([]byte("aaa\n"))
	w.Write([]byte("bbb\n"))
	w.Write([]byte("ccc\n"))
	fmt.Println(*c)
}

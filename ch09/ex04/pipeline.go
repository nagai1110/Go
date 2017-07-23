// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"time"
)

const (
	max = 1000000
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	fmt.Printf("[Start]%v \n", time.Now())
	go pipeLine(0, ch, done)

	ch <- 1
	<-done
}

func pipeLine(count int, ch <-chan int, done chan<- struct{}) {
	if count == max {
		fmt.Printf("[End]%v\nValue = %d\n", time.Now(), <-ch)
		done <- struct{}{}
	}

	count++
	nextCh := make(chan int)
	go pipeLine(count, nextCh, done)
	v := <-ch
	nextCh <- v
}

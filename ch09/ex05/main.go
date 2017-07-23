// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	go pingpong(ch1, ch2, done)
	go pingpong(ch2, ch1, done)

	ch1 <- 0
	ticker := time.Tick(1 * time.Second)
	<-ticker
	close(done)

	<-done

	var count int
	select {
	case count = <-ch1:
	case count = <-ch2:
	}
	fmt.Printf("Count = %d\n", count)
}

func pingpong(in <-chan int, out chan<- int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case count := <-in:
			out <- count + 1
		}
	}
}

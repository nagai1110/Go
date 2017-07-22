// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	channel chan<- string
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.channel <- "---Current members---"
			for client := range clients {
				cli.channel <- client.name
			}
			cli.channel <- "---------------------"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	const timeout = 5 * time.Minute

	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	closeFunc := func() {
		conn.Close()
	}
	ticker := time.AfterFunc(timeout, closeFunc)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		ticker.Stop()
		messages <- who + ": " + input.Text()
		ticker = time.AfterFunc(timeout, closeFunc)
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

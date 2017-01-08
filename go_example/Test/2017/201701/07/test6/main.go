package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	// messages := make(chan string, 1)
	signals := make(chan bool)

	messages <- "hello"
	fmt.Print("select1 -> ")
	select {
	case msg := <-messages:
		fmt.Println("received messages", msg)
	default:
		fmt.Println("no message received.")
	}

	msg := "hi"

	fmt.Print("select2 -> ")
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message send")
	}

	fmt.Print("select3 -> ")
	select {
	case msg := <-messages:
		fmt.Println("received message.", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no acitvity.")
	}
}

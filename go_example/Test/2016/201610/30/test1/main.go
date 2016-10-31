package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func producer(c chan bool) {
	defer close(c)
	time.Sleep(10 * time.Second)
	c <- true
	fmt.Println("发送消息")
}

func consumer(c chan bool) {

	cmd := exec.Command("ping", "127.0.0.1", "-t")
	cmd.Stdout = os.Stdout
	go cmd.Run()
	<-c
	cmd.Process.Kill()
	// 	// ok := true
	// 	// for ok {
	// 		if v, ok := <-c; ok {
	// 			time.Sleep(1 * time.Second)
	// 			fmt.Println("接受到消息", v)
	// 			goto Exit
	// 		}
	// 	}

}

func main() {
	c := make(chan bool)
	go producer(c)
	consumer(c)
}

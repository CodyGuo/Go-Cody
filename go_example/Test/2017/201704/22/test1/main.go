package main

import "fmt"

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type myTalk string

func (this *myTalk) Hello(userName string) string {
	return fmt.Sprint("hello ", userName)
}

func (this *myTalk) Talk(heard string) (saying string, end bool, err error) {
	return "talk: " + heard, true, nil
}

func main() {
	var talk Talk = new(myTalk)
	fmt.Println(talk.Hello("codyguo"))
}

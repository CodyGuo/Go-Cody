package main

import (
	"github.com/codyguo/logs"
)

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (this *fakeString) String() string {
	return this.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		logs.Notice(str)
	case Stringer:
		logs.Notice(str.String())
	default:
		logs.Error(str)
	}
}

func main() {
	s := &fakeString{"hello world."}
	printString("你好")
	printString(s)
}

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codyguo/logs"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			errMsg := fmt.Sprintf("recover --> %v", err)
			ioutil.WriteFile("error.log", []byte(errMsg), 0666)
			logs.Error(errMsg)
		}
	}()

	logs.Notice("hello world.")
	panic("panic error")
}

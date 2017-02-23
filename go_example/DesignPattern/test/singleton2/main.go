package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	SaySomething()
}

type singleton struct {
}

func (this singleton) SaySomething() {
	fmt.Println("Singleton.")
}

var singletonInstance Singleton
var once sync.Once

func NewSingletonInstance() Singleton {
	once.Do(func() {
		fmt.Println("newSingletonInstance.")
		singletonInstance = &singleton{}
	})
	return singletonInstance
}

func main() {
	sig := NewSingletonInstance()
	sig.SaySomething()

	sig = NewSingletonInstance()
	sig.SaySomething()
}

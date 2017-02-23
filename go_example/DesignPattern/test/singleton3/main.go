package main

import "fmt"

type Singleton interface {
	SaySomething()
}

type singleton struct {
}

func (this singleton) SaySomething() {
	fmt.Println("Singleton.")
}

var singletonInstance Singleton

func init() {
	fmt.Println("init new singleton.")
	singletonInstance = new(singleton)
}

func NewSingletonInstance() Singleton {
	return singletonInstance
}

func main() {
	sig := NewSingletonInstance()
	sig.SaySomething()
	sig = NewSingletonInstance()
	sig.SaySomething()
}

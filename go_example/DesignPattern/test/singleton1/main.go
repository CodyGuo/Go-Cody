/*
   单例模式 - nil
*/
package main

import "fmt"

type Singleton interface {
	SaySomething()
}

type singleton struct {
}

func (this singleton) SaySomething() {
	fmt.Println("Singleton")
}

var singletonInstance Singleton

func NewSingletonInstance() Singleton {
	if nil == singletonInstance {
		fmt.Println("first new...")
		singletonInstance = &singleton{}
	}
	return singletonInstance
}

func main() {
	// singletonInstance.SaySomething()
	sin := NewSingletonInstance()
	sin.SaySomething()
	sin = NewSingletonInstance()
	sin.SaySomething()
}

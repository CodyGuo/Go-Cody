package main

import (
	"fmt"
	"sync"
)

var m *Manager
var num int

func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
		fmt.Println("m...")
		num++
	}
	return m
}

type Manager struct{}

func (m Manager) Manage() {
	// fmt.Print("manage...\t")
}

func main() {
	goNum := 100000000
	wg := new(sync.WaitGroup)
	for i := 0; i < goNum; i++ {
		wg.Add(1)
		go func() {
			m = GetInstance()
			m.Manage()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}

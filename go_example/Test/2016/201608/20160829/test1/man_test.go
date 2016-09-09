package main

import (
	"testing"
)

func ExampleRec() {
	recPrint(5)
}

func TestRec(t *testing.T) {
	recPrint(100)
}

func BenchmarkRec(t *testing.B) {
	for i := 0; i < t.N; i++ {
		recPrint(i)
	}
}

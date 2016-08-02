package main

import (
	"testing"
)

func BenchmarkMaxNumstr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNumStr(0)
		maxNumStr(1)
		maxNumStr(8)
		maxNumStr(10)
		maxNumStr(100)
		maxNumStr(1000)
	}
}

func BenchmarkMaxNumPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNumPow(0)
		maxNumPow(1)
		maxNumPow(8)
		maxNumPow(10)
		maxNumPow(100)
		maxNumPow(1000)
	}
}

func BenchmarkMaxNumAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNumPow(0)
		maxNumAdd(1)
		maxNumAdd(8)
		maxNumAdd(10)
		maxNumAdd(100)
		maxNumAdd(1000)
	}
}

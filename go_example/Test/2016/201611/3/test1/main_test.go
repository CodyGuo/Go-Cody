package main

import (
	"testing"
)

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switchTest(0)
		switchTest(10)
		switchTest(100)
		switchTest(1000)
	}
}

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifTest(0)
		ifTest(10)
		ifTest(100)
		ifTest(1000)
	}
}

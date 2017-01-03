package main

import (
	"testing"
)

var (
	nums = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
)

func BenchmarkInshort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inshort(nums)
	}
}

func BenchmarkShshrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shshort(nums)
	}
}

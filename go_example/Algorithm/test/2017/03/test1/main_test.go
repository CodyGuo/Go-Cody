package main

import (
	"testing"
)

var (
	items = []int{1, 3, 5, 7, 9, 10, 15, 16, 18, 20, 22, 23, 28, 29, 35, 50, 52, 66, 68, 69, 71, 72, 75, 78, 79, 89, 91, 99, 100, 120}
)

func BenchmarkBinarySearch1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearch(items, 10)
		binarySearch(items, 1000)
		binarySearch(items, 100)
		binarySearch(items, 5)
	}
}

func BenchmarkBinarySearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearch2(items, 10)
		binarySearch2(items, 1000)
		binarySearch2(items, 100)
		binarySearch2(items, 5)
	}
}

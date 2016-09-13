package main

import (
	"testing"
)

var (
	nums []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
)

func BenchmarkInsertionSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort1(nums)
	}
}

func BenchmarkInsertionSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort2(nums)
	}
}

func BenchmarkInsertionSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort3(nums)
	}
}

func BenchmarkInsertionSort4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort4(nums)
	}
}

func BenchmarkInsertionSort5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort5(nums)
	}
}

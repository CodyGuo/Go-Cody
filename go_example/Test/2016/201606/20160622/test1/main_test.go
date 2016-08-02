package main

import (
	"testing"
)

func BenchmarkEasySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		easySum(100)
		easySum(1000)
		easySum(10000)
		easySum(100000)
	}
}

func BenchmarkGaosiSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gaosiSum1(100)
		gaosiSum1(1000)
		gaosiSum1(10000)
		gaosiSum1(100000)
	}
}

func BenchmarkGaosiSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gaosiSum2(100)
		gaosiSum2(1000)
		gaosiSum2(10000)
		gaosiSum2(100000)
	}
}

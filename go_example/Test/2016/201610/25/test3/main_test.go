package main

import (
	"testing"
)

func Benchmark_InvSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InvSqrt(65535)
		InvSqrt(65535323232)
	}
}

func Benchmark_mathSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathSqrt(65535)
		mathSqrt(65535323232)
	}
}

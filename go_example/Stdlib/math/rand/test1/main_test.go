package main

import (
	"testing"
)

func BenchmarkGetNonce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNonce1()
	}
}

func BenchmarkGetNonce2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNonce2()
	}
}

func BenchmarkGetNonce3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNonce3(24)
	}
}

func BenchmarkToString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString1()
	}
}

func BenchmarkTostring2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString2()
	}
}

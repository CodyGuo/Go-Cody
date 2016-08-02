package main

import (
	"testing"
)

var (
	str1 = "123164987321313496797979879645413416346"
	str2 = "123"
	str3 = "12"
	str4 = ""
)

func BenchmarkComma1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma1(str1)
		comma1(str2)
		comma1(str3)
		comma1(str4)
	}
}

func BenchmarkComma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma2(str1)
		comma2(str2)
		comma2(str3)
		comma2(str4)
	}
}

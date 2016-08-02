package main

import (
	"testing"
)

var (
	nums1 = []int{1, 2, 3, 4, 5, 6, 6, 7}
	nums2 = []int{0}
	nums3 = []int{2, 3, 4, 56, 6234, 234, 6234, 234234, 234234, 665433, 4, 56, 6234, 234, 6234, 234234, 3, 4, 56, 6234, 234, 6234, 234234}
)

func BenchmarkAvg1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		average1(nums1)
		average1(nums2)
		average1(nums3)
	}
}

func BenchmarkAvg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		average2(nums1)
		average2(nums2)
		average2(nums3)
	}
}

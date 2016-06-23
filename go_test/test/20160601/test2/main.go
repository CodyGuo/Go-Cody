package main

import (
	"fmt"
)

// 计算两个int切片类型的交集，不含重复元素
func IntersectNew1(a, b []int) []int {
	count1 := len(a)
	if count1 == 0 {
		return nil
	}
	count2 := len(b)
	if count2 == 0 {
		return nil
	}
	m := map[int]bool{}
	if count1 < count2 {
		c := make([]int, 0, count1)
		for i := 0; i < count1; i++ {
			m[a[i]] = true
		}
		for _, x := range b {
			if m[x] {
				l := len(c)
				c = c[:l+1]
				c[l] = x
				delete(m, x)
			}
		}
		return c
	} else {
		c := make([]int, 0, count2)
		for i := 0; i < count2; i++ {
			m[b[i]] = true
		}
		for _, x := range a {
			if m[x] {
				l := len(c)
				c = c[:l+1]
				c[l] = x
				delete(m, x)
			}
		}
		return c
	}
	m = nil //内存回收
	return nil
}

// 计算两个int切片类型的交集，不含重复元素
func IntersectNew2(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	m := map[int]bool{}
	if len(b) < len(a) {
		c := make([]int, 0, len(a))
		for i := 0; i < len(a); i++ {
			m[a[i]] = true
		}
		for _, x := range b {
			if m[x] {
				l := len(c)
				c = c[:l+1]
				c[l] = x
				delete(m, x)
			}
		}
		return c
	} else {
		c := make([]int, 0, len(b))
		for i := 0; i < len(b); i++ {
			m[b[i]] = true
		}
		for _, x := range a {
			if m[x] {
				l := len(c)
				c = c[:l+1]
				c[l] = x
				delete(m, x)
			}
		}
		return c
	}
	m = nil //内存回收
	return nil
}

// 计算两个int切片类型的交集，不含重复元素
func Intersect1(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	m := map[int]bool{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = true
	}

	c := make([]int, 0, len(m)+1)
	for _, x := range b {
		if m[x] {
			c = append(c, x)
			delete(m, x)
		}
	}
	// m = nil //内存回收
	return c
}

// 计算两个int切片类型的交集，不含重复元素
func Intersect2(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	c := []int{}

	m := map[int]bool{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = true
	}

	for _, x := range b {
		if m[x] {
			c = append(c, x)
			// m[x] = false
			delete(m, x)
		}
	}
	// m = nil //内存回收
	return c
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	tmp := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num1 == num2 {
				tmp[num1] = num1
			}
		}
	}

	result := []int{}
	for _, v := range tmp {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3, 1}, []int{2, 3}))
}

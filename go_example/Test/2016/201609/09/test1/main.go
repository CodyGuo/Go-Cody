package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	num := []int{432, 432432, 0, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	insertionSort1(num)
	fmt.Println(num)
}

// 插入排序
func insertionSort1(nums []int) {
	for j := 1; j < len(nums); j++ {
		// 增加判断减少循环
		if nums[j] < nums[j-1] {
			key := nums[j]

			i := j - 1
			for ; i >= 0 && nums[i] > key; i-- {
				nums[i+1] = nums[i]
			}
			nums[i+1] = key
		}

	}
}

func insertionSort2(old []int) {
	insertionSort(old)
}

func insertionSort(old []int) (sortedData *list.List, err error) {
	sortedData = list.New()
	sortedData.PushBack(old[0])
	size := len(old)
	for i := 1; i < size; i++ {
		v := old[i]
		e := sortedData.Front()
		for nil != e {
			if e.Value.(int) >= v {
				sortedData.InsertBefore(v, e)
				break
			}
			e = e.Next()
		}
		//the biggest,put @v on the back of the list
		if nil == e {
			sortedData.PushBack(v)
		}
	}

	return
}

func insertionSort3(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
}

func insertionSort4(nums []int) {
	sort.Ints(nums)
}

func InsertionSort5(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			swap(nums, j, j-1)
		}
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

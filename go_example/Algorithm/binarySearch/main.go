package main

import "sort"

// 右移1位
func binarySearch(items []int, item int) int {
	result := -1
	low, high := 0, len(items)-1
	for low <= high {
		mid := (low + high) >> 1
		if items[mid] < item {
			low = mid + 1
		} else if items[mid] > item {
			high = mid - 1
		} else {
			result = mid
			break
		}
	}
	return result
}

// 除2
func binarySearch2(items []int, item int) int {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := items[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func binarySearchStd(items []int, item int) int {
	return sort.SearchInts(items, item)
}

func binarySearch3(items []int, item int) int {
	det := -1
	left, right := 0, len(items)-1
	for left <= right {
		mid := (left + right) >> 1
		if items[mid] < item {
			left = mid + 1
		} else if items[mid] > item {
			right = mid - 1
		} else {
			det = mid
			break
		}
	}
	return det
}

func main() {
	items := []int{1, 3, 5, 7, 9, 10, 15, 16, 18, 20, 22, 23, 28, 29, 35, 50, 52, 66}
	println(binarySearch(items, 3))
	println(binarySearch(items, -1))
	println(binarySearch(items, 7))
	println(binarySearch(items, 10))
	println(binarySearch(items, 22))
	println(binarySearch(items, 60))
	println(binarySearch(items, 50))
}

package main

// 右移1位
func binarySearch(items []int, item int) int {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) >> 1
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

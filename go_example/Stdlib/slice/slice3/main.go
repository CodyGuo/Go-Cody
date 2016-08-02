package main

import (
	"fmt"
)

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	var nums = []int{2, 1, 5, 3, 15, 8}
	fmt.Println(nums)
	sort(nums)

	fmt.Println(nums)
}

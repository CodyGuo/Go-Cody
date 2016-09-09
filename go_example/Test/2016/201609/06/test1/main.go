package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 4, 3, 0}
	target := 0

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, n := range nums {
		for j, m := range nums[i+1:] {
			// fmt.Println(nums[i+1+j:])
			if target == n+m {
				// fmt.Printf("[%d] [%d] %d + %d = %d\n ", i, j, n, m, target)
				result = append(result, i, i+1+j)
			}
		}
	}

	return result
}

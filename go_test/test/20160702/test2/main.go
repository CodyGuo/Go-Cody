package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(average1(nums))
	fmt.Println(average2(nums))

}

func average1(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	return sum / len(nums)
}

func average2(nums []int) int {
	var avg int = 1
	N := len(nums)
	sum := avg * N
	sum = sum - avg + nums[1]
	return sum / N
}

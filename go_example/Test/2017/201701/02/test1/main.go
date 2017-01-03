package main

import "fmt"

func inshort(nums []int) {
	for i := 1; i < len(nums); i++ {
		lookout := nums[i]
		j := i - 1
		for ; j >= 0 && lookout < nums[j]; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = lookout
	}
}

func shshort(nums []int) {
	n := len(nums)
	d := n / 2
	for d >= 1 {
		for i := d + 1; i >= 0 && i < n; i++ {
			lookout := nums[i]
			j := i - d
			for ; j >= 0 && lookout < nums[j]; j -= d {
				nums[j+d] = nums[j]
			}
			nums[j+d] = lookout
		}
		d = d / 2
	}
}

func main() {
	var nums = []int{432, 432432, 0, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	fmt.Printf("排序前: \n%v\n", nums)
	// inshort(nums)
	shshort(nums)
	fmt.Println(nums)
}

package main

import "fmt"

func bubble(nums []int) {
	size := len(nums)
	for {
		swapped := false
		for i := range size - 1 {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubble(arr)

	fmt.Println(arr)
}

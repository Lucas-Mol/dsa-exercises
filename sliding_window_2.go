package main

func maxSumSubarray(nums []int, maxRange int) (maxSum int) {
	if len(nums) < 0 {
		return
	}

	windowSum := 0
	for i := 0; i < maxRange; i++ {
		windowSum += nums[i]
	}

	maxSum = windowSum
	for i := maxRange; i < len(nums); i++ {
		windowSum = windowSum - nums[i-maxRange] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	println("Max Sum:", maxSumSubarray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4))
}

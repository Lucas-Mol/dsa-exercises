package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	seen := make(map[int]struct{})

	for i, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}

		if len(seen) > k {
			delete(seen, nums[i-k])
		}
	}

	return false
}

func main() {
	println("Contains:", containsNearbyDuplicate([]int{1, 2, 3, 1, 5, 6, 7, 8, 9, 10}, 4))
}

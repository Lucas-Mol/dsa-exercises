package main

func binarySearch(arr []int, target int) (index int) {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return -1
}

func main() {
	println("Index:", binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}

package main

func exponentialSearch(arr []int, target int) (index int) {
	if len(arr) == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	n := len(arr)
	i := 1

	for i < n && arr[i] < target {
		i *= 2
	}

	if i > n {
		i = n - 1
	}

	if arr[i] == target {
		return i
	}

	start := i / 2
	end := i
	if n-1 < end {
		end = n - 1
	}
	return _binarySearch(arr, target, start, end)
}

func _binarySearch(arr []int, target int, low int, high int) (index int) {
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
	println("Index:", exponentialSearch([]int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18}, 16))
}

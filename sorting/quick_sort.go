package main

import "fmt"

func recursiveQuick(a *[]int, left int, right int) {
	if left < right {
		pi := partition(a, left, right)
		recursiveQuick(a, left, pi-1)
		recursiveQuick(a, pi+1, right)
	}
}

func partition(a *[]int, left int, right int) int {
	pivot := (*a)[right] //dumb pivot choice
	i := left - 1

	for j := left; j < right; j++ {
		if (*a)[j] <= pivot {
			i++
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		}
	}

	(*a)[i+1], (*a)[right] = (*a)[right], (*a)[i+1]
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	recursiveQuick(&arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}

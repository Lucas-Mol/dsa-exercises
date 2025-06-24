package main

func minimumSwaps(arr []int32) int32 {
	var swaps int32
	for i := 0; i < len(arr); {
		if arr[i] != int32(i+1) {
			correctIndex := arr[i] - 1
			arr[i], arr[correctIndex] = arr[correctIndex], arr[i]
			swaps++
		} else {
			i++
		}
	}
	return swaps
}

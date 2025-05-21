package main

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	for i, r := range s {
		if counter[r] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	println(firstUniqChar("leetcode"))
}

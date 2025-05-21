package main

func maxLengthSubstrings(s string, maxOccurrences int) (max int) {
	if len(s) == 0 {
		return
	}

	l, r := 0, 0
	max = 1
	counter := make(map[rune]int)
	counter[rune(s[0])] = 1

	for r < len(s)-1 {
		r += 1
		if _, ok := counter[rune(s[r])]; ok {
			counter[rune(s[r])]++
		} else {
			counter[rune(s[r])] = 1
		}
		for counter[rune(s[r])] > maxOccurrences {
			counter[rune(s[l])]--
			l += 1
		}

		currentLength := r - l + 1
		if currentLength > max {
			max = currentLength
		}
	}

	return max
}

func main() {
	println("Max:", maxLengthSubstrings("bcbbbcbadfghdb", 2))
}

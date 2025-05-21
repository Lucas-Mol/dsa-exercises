package main

//Not using any lib

func reverseWord(word string) string {
	runes := []rune(word)

	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

func splitIntoWords(sentence string) (words []string) {
	for l, r := 0, 0; r < len(sentence); r++ {
		if sentence[r] == ' ' || r == len(sentence)-1 {
			if r == len(sentence)-1 {
				r++
			}
			words = append(words, sentence[l:r])
			l = r + 1
		}
	}
	return words
}

func joinWords(words []string, separator string) (result string) {
	for _, word := range words {
		result += word + separator
	}
	return result[:len(result)-1]
}

func ReverseEachWord(sentence string) string {
	words := splitIntoWords(sentence)
	for i, word := range words {
		println(i, word)
		words[i] = reverseWord(word)
	}
	return joinWords(words, " ")
}

func main() {
	println(ReverseEachWord("The quick brown fox jumps over the lazy dog"))
}

package valid_anagram

func isAnagram(s string, t string) bool {
	charMap := make(map[byte]int)
	for _, char := range []byte(s) {
		_, ok := charMap[char]
		if ok {
			charMap[char]++
		} else {
			charMap[char] = 1
		}
	}
	for _, char := range []byte(t) {
		value, ok := charMap[char]
		if ok {
			if value > 1 {
				charMap[char]--
			} else {
				delete(charMap, char)
			}
		} else {
			return false
		}
	}

	return len(charMap) == 0
}

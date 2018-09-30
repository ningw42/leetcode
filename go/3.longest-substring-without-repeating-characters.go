func lengthOfLongestSubstring(s string) int {
	var target, current string
	bucket := make(map[byte]int)
	target = ""

outer:
	for begin := 0; true; begin++ {
		current = ""
		if begin == len(s) {
			break
		}
		for i := begin; true; i++ {
			if i == len(s) {
				if len(current) > len(target) {
					target = current
				}
				break outer
			}
			char := s[i]
			if _, duplicate := bucket[char]; duplicate {
				bucket = make(map[byte]int)
				break
			}
			current += string(char)
			bucket[char] = 1
		}
		if len(current) > len(target) {
			target = current
		}
	}

	return len(target)
}
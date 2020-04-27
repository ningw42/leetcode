func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[1]) + s
		}
	}
	var lefts, rights []string
	for i := 0; i < len(s) / 2; i++ {
		lefts = append(lefts, reverseString(s[:i+1]))
		rights = append(rights, s[i+1:])
	}

	var left, right string
	for i := len(lefts)-1; i >= 0; i-- {
		left = lefts[i]
		right = rights[i]
		if len(right) > len(left) && left == right[1:1+len(left)] {
			return reverseString(right[1:]) + right
		}
		if left == right[:len(left)] {
			return reverseString(right) + right
		}
	}
	return reverseString(s[1:]) + s
}

func reverseString(s string) string {
	bs := []byte(s)
	for i := 0; i < len(s) / 2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	return string(bs)
}


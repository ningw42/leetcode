func findWords(words []string) []string {
	row := [][]byte{
		[]byte("qwertyuiop"),
		[]byte("asdfghjkl"),
		[]byte("zxcvbnm"),
	}

	m := make(map[byte]int)
	for i, chars := range row {
		for _, char := range chars {
			m[char] = i
		}
	}

	var ret []string
	for _, word := range words {
		sameRow := true
		chars := bytes.ToLower([]byte(word))
		rowIndex := m[chars[0]]
		for _, char := range chars[1:] {
			if rowIndex != m[char] {
				sameRow = false
				break
			}
		}
		if sameRow {
			ret = append(ret, word)
		}
	}

	return ret
}

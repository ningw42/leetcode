func titleToNumber(s string) int {
	row := 0
	for i := 0; i < len(s); i++ {
		digit := s[i]
		row = row * 26 + int(digit - 'A') + 1
	}
	return row
}
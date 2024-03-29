func lengthOfLastWord(s string) int {
	begin := -1
	end := math.MaxInt64
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	if end == math.MaxInt64 {
		return 0
	}
	for ; i >= 0; i-- {
		if s[i] == ' ' {
			begin = i
			break
		}
	}

	return end - begin
}
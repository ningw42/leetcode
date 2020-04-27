func reverseWords(s string) string {
	var ss []string
	begin := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if begin != -1 {
				ss = append(ss, s[begin:i])
				begin = -1
			}
		} else {
			if begin == -1 {
				begin = i
			}
		}
	}
	if begin != -1 {
		ss = append(ss, s[begin:])
	}

	for i := 0; i < len(ss) / 2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}

	return strings.Join(ss, " ")
}
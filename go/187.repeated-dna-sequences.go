func findRepeatedDnaSequences(s string) []string {
	count := make(map[string]int)
	var repeated []string
	for i := 0; i+10 <= len(s); i++ {
		substring := s[i : i+10]
		if c, e := count[substring]; e && c == 1 {
			repeated = append(repeated, substring)
		}
		count[substring]++
	}

	return repeated
}

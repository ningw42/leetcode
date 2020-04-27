func numJewelsInStones(J string, S string) int {
	stones := make(map[byte]int)
	ret := 0

	for _, r := range S {
		c := byte(r)
		if count, exists := stones[c]; exists {
			stones[c] = count + 1
		} else {
			stones[c] = 1
		}
	}
    
	for _, r := range J {
		c := byte(r)
		ret += stones[c]
	}

	return ret
}
func toLowerCase(str string) string {
	var ret []byte

	for _, r := range str {
		c := byte(r)
		if c >= 65 && c <= 90 {
			c = c + 32
		}
		ret = append(ret, c)
	}

	return string(ret)
}
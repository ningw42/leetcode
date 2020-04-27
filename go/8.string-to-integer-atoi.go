func myAtoi(str string) int {
	var chars []byte
	var minus bool
	str = strings.TrimLeft(str, " ")

	if len(str) == 0 {
		return 0
	}
	if str[0] == '-' {
		minus = true
		str = str[1:]
	} else if str[0] == '+' {
		minus = false
		str = str[1:]
	} else if '0' <= str[0] && str[0] <= '9' {
		minus = false
	} else {
		return 0
	}

	for _, r := range str {
		char := byte(r)
		if char == '0' {
			if chars != nil {
				chars = append(chars, char)
			}
		} else if '1' <= char && char <= '9' {
			chars = append(chars, char)
		} else {
			break
		}
	}

	// prevent from int64 overflow
	if len(chars) > 10 {
		if minus {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	var sum int
	weight := 1
	for i := len(chars) - 1; i >= 0; i-- {
		sum += int(chars[i] - '0') * weight
		weight = weight * 10
	}

	if minus && sum > math.MaxInt32 + 1 {
		return math.MinInt32
	}
	if !minus && sum > math.MaxInt32 {
		return math.MaxInt32
	}

	if minus {
		return -sum
	} else {
		return sum
	}
}
func divide(dividend int, divisor int) int {
	var minus bool
	if dividend < 0 {
		dividend = -dividend
		minus = !minus
	}
	if divisor < 0 {
		divisor = -divisor
		minus = !minus
	}

	var q int
	for dividend >= divisor {
		dividend -= divisor
		q++
	}

	if minus {
		if q < math.MinInt32 {
			return math.MaxInt32
		} else {
			return -q
		}
	} else {
		if q > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return q
		}
	}
}
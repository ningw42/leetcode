func myPow(x float64, n int) float64 {
	var minusN, minusX bool
	if x < 0 {
		x = -x
		minusX = true
	}
	if n < 0 {
		n = -n
		minusN = true
	}

	pow := naivePow(x, n)

	var result float64
	if minusN {
		result = 1.0 / pow
	} else {
		result = pow
	}
	if minusX && n % 2 == 1 {
		result = -result
	}
	return result
}

func naivePow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		half := naivePow(x, n / 2)
		if result := half * half; n % 2 == 0 {
			return result
		} else {
			return result * x
		}
	}
}
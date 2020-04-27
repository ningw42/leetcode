func mySqrt(x int) int {
	begin := 0
	end := x
	for {
		mid := (begin + end) / 2
		if sqr := mid * mid; sqr > x {
			end = mid
		} else if sqr == x {
			return mid
		} else {
			if sqr2 := (mid+1) * (mid+1); sqr2 > x {
				return mid
			} else if sqr2 == x {
				return mid + 1
			} else {
				begin = mid
			}
		}
	}
}
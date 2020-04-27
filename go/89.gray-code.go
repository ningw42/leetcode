func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	} else {
		prev := grayCode(n-1)
		ret := make([]int, 2*len(prev))
		for i := 0; i < len(prev); i++ {
			ret[i] = prev[i] * 2
			ret[2*len(prev)-1-i] = prev[i] * 2 + 1
		}
		return ret
	}
}
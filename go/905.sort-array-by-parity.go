func sortArrayByParity(A []int) []int {
	odd := []int{}
	even := []int{}

	for _, item := range A {
		if item & 1 == 1{
			odd = append(odd, item)
		} else {
			even = append(even, item)
		}
	}

	return append(even, odd...)
}
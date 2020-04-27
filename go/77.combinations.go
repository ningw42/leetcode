
var result [][]int
func combine(n int, k int) [][]int {
	result = [][]int{}
	findCombination(1, n, k, []int{})
	return result
}

func findCombination(begin int, n int, k int, combination []int) {
	if k == 0 {
		result = append(result, makeSliceCopy(combination))
	} else {
		for i := begin; i <= n; i++ {
			findCombination(i+1, n, k-1, append(combination, i))
		}
	}
}

func makeSliceCopy(s []int) []int {
	copy := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		copy[i] = s[i]
	}
	return copy
}
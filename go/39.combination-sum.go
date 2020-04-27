
var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = [][]int{}
	findCombination([]int{}, candidates, target)

	return result
}

func findCombination(combination []int, candidates []int, target int) {
	if target == 0 {
		result = append(result, makeSliceCopy(combination))
	} else if target < 0 {
		return 
	} else {
		for i := 0; i < len(candidates); i++ {
			findCombination(append(combination, candidates[i]), candidates[i:], target - candidates[i])
		}
	}
}

func makeSliceCopy(s []int) []int {
	copy := make([]int, len(s))
	for i, e := range s {
		copy[i] = e
	}
	return copy
}
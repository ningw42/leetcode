
var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
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
		for i := 0; i < len(candidates); {
			findCombination(append(combination, candidates[i]), candidates[i+1:], target - candidates[i])
			for i = i + 1; i < len(candidates) && candidates[i] == candidates[i-1]; i++ {}
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
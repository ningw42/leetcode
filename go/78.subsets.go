func subsets(nums []int) [][]int {
	var subsets [][]int
	var subset []int
	var index int
	num := int(math.Pow(2, float64(len(nums))))

	for i := 0; i < num; i++ {
		index = 0
		subset = []int{}
		for j := i; j != 0; j = j / 2 {
			if j%2 == 1 {
				subset = append(subset, nums[index])
			}
			index++
		}
		subsets = append(subsets, subset)
	}

	return subsets
}

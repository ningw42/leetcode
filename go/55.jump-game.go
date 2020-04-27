func canJump(nums []int) bool {
	steps := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		step := math.MaxInt32
		for j := nums[i]; j >= 1; j-- {
			if i+j >= len(nums) {
				step = 0
				break
			}
			if newStep := steps[i+j]; newStep < step {
				step = newStep
			}
		}
		steps[i] = step + 1
	}
	return steps[0] != math.MaxInt32 + 1
}
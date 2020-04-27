func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	count := 0
	i := 1
	for i < length {
		if nums[i] == nums[i-1] {
			count++
			if count == 2 {
				count--
				length--
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				i++
			}
		} else {
			count = 0
			i++
		}
	}
	return length
}
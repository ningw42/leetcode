func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	c1 := helper(nums[1:])
	c2 := helper(nums[0 : len(nums)-1])
	if c1 > c2 {
		return c1
	} else {
		return c2
	}
}

func helper(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	results := make([]int, len(nums))
	results[0] = nums[0]
	if nums[0] > nums[1] {
		results[1] = nums[0]
	} else {
		results[1] = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		c1 := results[i-2] + nums[i]
		c2 := results[i-1]
		if c1 > c2 {
			results[i] = c1
		} else {
			results[i] = c2
		}
	}

	return results[len(nums)-1]
}

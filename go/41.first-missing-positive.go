func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] < 0 || nums[i] >= length {
			nums[i] = 0
		}
	}
	for i := 0; i < length; i++ {
		nums[nums[i]%length] += length
	}
	for i := 1; i < length; i++ {
		if nums[i] < length {
			return i
		}
	}

	return length 
}

// O(n) time complexity but O(n) space complexity
func naive(nums []int) int {
	all := make(map[int]bool)
	for _, num := range nums {
		if num > 0 {
			all[num] = true
		}
	}

	i := 1
	for {
		if _, exists := all[i]; !exists {
			break
		}
		i++
	}

	return i
}
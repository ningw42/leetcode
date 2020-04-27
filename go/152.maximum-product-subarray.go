func maxProduct(nums []int) int {
	var product int
	max := math.MinInt32
	products := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if j == i {
				product = nums[i]
			} else {
				product = products[j-1] * nums[j]
			}
			products[j] = product
			if product > max {
				max = product
			}
		}
	}

	return max
}

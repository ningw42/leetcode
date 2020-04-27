func singleNumber(nums []int) int {
	return withExtraMemory(nums)    
}

func withExtraMemory(nums []int) int {
	dist := make(map[int]int)
	for _, num := range nums {
	    dist[num]++
	}

	for num, c := range dist {
		if c == 1 {
			return num
		}
	}
	return 0
}

func withoutExtraMemory(nums []int) int {
	return 0
}
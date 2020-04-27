func longestConsecutive(nums []int) int {
	sequence := make(map[int][]int)
	done := make(map[int]bool)
	for _, num := range nums {
		sequence[num] = []int{num, num}
		done[num] = false
	}
	for num := range sequence {
		if visited := done[num]; visited {
			continue
		}
		var cursor, left, right int
		var exists bool

		// go left
		cursor = num - 1
		_, exists = sequence[cursor]
		for exists {
			done[cursor] = true
			cursor--
			_, exists = sequence[cursor]
		}
		left = cursor + 1
		// go right
		cursor = num + 1
		_, exists = sequence[cursor]
		for exists {
			done[cursor] = true
			cursor++
			_, exists = sequence[cursor]
		}
		right = cursor - 1

		// write result
		sequence[num] = []int{left, right}
	}
	// fmt.Println(sequence)
	maxLength := 0
	for _, interval := range sequence {
		if length := interval[1] - interval[0] + 1; length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
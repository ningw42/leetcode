func candy(ratings []int) int {
	// 1. find all local miniums(less or equals to both side)
	// 2. for each local minium, 
	//    increase number of candies from the minium to both side 
	//    until meets a local peak
	if len(ratings) <= 1 {
		return len(ratings)
	}
	var minIndexes []int
	for i := 0; i < len(ratings); i++ {
		if isLocalMinium(ratings, i) {
			minIndexes = append(minIndexes, i)
		}
	}

	candies := make([]int, len(ratings))
	for _, minIndex := range minIndexes {
		var cursor, currentMinCandyCount int
		candies[minIndex] = 1
		// go left
		cursor = minIndex - 1
		currentMinCandyCount = 1
		for cursor >= 0 {
			if ratings[cursor] > ratings[cursor+1] {
				currentMinCandyCount++
			} else {
				break
			}
			if candies[cursor] < currentMinCandyCount {
				candies[cursor] = currentMinCandyCount
			}
			cursor--
		}
		// go right
		cursor = minIndex + 1
		currentMinCandyCount = 1
		for cursor < len(ratings) {
			if ratings[cursor] > ratings[cursor-1] {
				currentMinCandyCount++
			} else {
				break
			}
			if candies[cursor] < currentMinCandyCount {
				candies[cursor] = currentMinCandyCount
			}
			cursor++
		}
	}

	var sum int
	for _, candy := range candies {
		sum += candy
	}
	return sum
}

func isLocalMinium(ratings []int, i int) bool {
	if i == 0 {
		return ratings[i] <= ratings[1]
	}
	if i == len(ratings) - 1 {
		return ratings[i] <= ratings[i-1]
	}
	return ratings[i-1] >= ratings[i] && ratings[i] <= ratings[i+1]
}
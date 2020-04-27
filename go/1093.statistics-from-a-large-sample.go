func sampleStats(count []int) []float64 {
	var min, max, mean, median, mode float64
	var total, sum, maxCount int
	maxCount = 0

	// mean
	for i, c := range count {
		total += c
		sum += i * c
		if c > maxCount {
			maxCount = c
			mode = float64(i)
		}
	}
	mean = float64(sum) / float64(total)
	// min
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			min = float64(i)
			break
		}
	}
	// max
	for i := len(count) - 1; i >= 0; i-- {
		if count[i] != 0 {
			max = float64(i)
			break
		}
	}
	// median
	var medianIndex, medians []int
	if total&1 == 1 {
		medianIndex = []int{total / 2}
	} else {
		medianIndex = []int{total/2 - 1, total / 2}
	}
	numberOfLeftElement := 0
	for i, c := range count {
		numberOfLeftElement += c
		for ii := 0; ii < len(medianIndex); {
			if numberOfLeftElement >= medianIndex[ii]+1 {
				medians = append(medians, i)
				medianIndex = medianIndex[1:]
			} else {
				ii++
			}
		}
		if len(medianIndex) == 0 {
			break
		}
	}
	medianSum := 0
	for _, med := range medians {
		medianSum += med
	}
	fmt.Println(medians)
	median = float64(medianSum) / float64(len(medians))

	return []float64{min, max, mean, median, mode}
}

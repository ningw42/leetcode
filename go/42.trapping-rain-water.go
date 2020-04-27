func trap(height []int) int {
	// find local peak
	var peak []int
	var leftHeight, rightHeight int
	for i := 0; i < len(height); i++ {
		if i == 0 {
			leftHeight = 0
		} else {
			leftHeight = height[i-1]
		}
		if i == len(height) - 1 {
			rightHeight = 0
		} else {
			rightHeight = height[i+1]
		}
		if h := height[i]; h >= leftHeight && h >= rightHeight {
			peak = append(peak, i)
		}
	}
	if len(peak) <= 1 {
		return 0
	}

	fmt.Println(peak)
	// remove non-edge local peak
	var removed bool
	for {
		removed = false
		for i := 1; i < len(peak)-1; i++ {
			if height[peak[i]] <= height[peak[i-1]] && height[peak[i]] <= height[peak[i+1]] {
				peak = append(peak[:i], peak[i+1:]...)
				removed = true
				break
			}
		}
		if !removed {
			break
		}
	}
	fmt.Println(peak)

	// add up 
	var water, left, right, localHeight int
	for i := 1; i < len(peak); i++ {
		left = peak[i-1]
		right = peak[i]
		localHeight = min(height[left], height[right])
		for j := left+1; j < right; j++ {
			if height[j] < localHeight {
				water += localHeight - height[j]
			}
		}
	}
	
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
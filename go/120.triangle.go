/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	result := [][]int{[]int{triangle[0][0]}}
	for i := 1; i < len(triangle); i++ {
		prev := result[i-1]
		var min []int
		for j, w := range triangle[i] {
			if j - 1 < 0 {
				min = append(min, w + prev[j])
			} else if j >= len(prev) {
				min = append(min, w + prev[j-1])
			} else {
				left := w + prev[j-1]
				right := w + prev[j]
				if left > right {
					min = append(min, right)
				} else {
					min = append(min, left)
				}
			}
		}
		result = append(result, min)
	}

	min := math.MaxInt32
	for _, sum := range result[len(result)-1] {
		if sum < min {
			min = sum
		}
	}

	return min
}


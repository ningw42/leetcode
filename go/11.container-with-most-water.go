/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(heights []int) int {
    i := 0;
    j := len(heights) - 1
    max := 0
    for i < j {
        if area := (j-i) * min(heights[i], heights[j]); area > max {
            max = area
        }
        if heights[i] < heights[j] {
            for i < j && heights[i+1] <= heights[i] {i++}
            i++
        } else if heights[i] >= heights[j] {
            for i < j && heights[j-1] <= heights[j] {j--}
            j--
        }
    }
    return max
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}


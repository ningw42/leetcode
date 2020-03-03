/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	return MonotoneStack(heights)
}

func MonotoneStack(heights []int) int {
	left := make([]int, len(heights))
    right := make([]int, len(heights))
    var stack []int
    for i, h := range heights {
        for len(stack) != 0 && heights[stack[len(stack)-1]] >= h {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            left[i] = -1
        } else {
            left[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
    stack = nil
    for i := len(heights)-1; i >= 0; i-- {
        h := heights[i]
        for len(stack) != 0 && heights[stack[len(stack)-1]] >= h {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            right[i] = len(heights)
        } else {
            right[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
 
    max := 0
    for i, h := range heights {
        if area := h * (right[i]-left[i]-1); area > max {
            max = area
        }
    }
    
    return max
}

func BruteForce(heights []int) int {
	max := 0
    for i := 0; i < len(heights); i++ {
		height := heights[i]
		for j := i; j >= 0; j-- {
			if heights[j] < height {
				height = heights[j]
			}
			area := height * (i - j + 1)
			if max < area {
				max = area
			}
		}
	}

	return max
}


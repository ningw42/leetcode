/*
 * @lc app=leetcode id=795 lang=golang
 *
 * [795] Number of Subarrays with Bounded Maximum
 */

// @lc code=start
func numSubarrayBoundedMax(A []int, L int, R int) int {
	return LinearSearch(A, L, R)
}

func LinearSearch(A []int, L, R int) int {
	begin, end, count := -1, -1, 0
	for i, a := range A {
		if a < L {
			// if an element E is smaller than L,
			// all the subarrays ends at E and starting from
			// the last element that greater than R and contains
			// the last element that fits in the range should be count
			// which is exactly "end - begin"
		}
		if L <= a && a <= R {
			// if an element E fits in the range,
			// all the subarrays ends at E starting from the last
			// element that greater than R should be count
			end = i
		}
		if a > R {
			// if an element E is greater than R,
			// new segmentation begins
			begin = i
			end = i
		}

		count += end - begin
	}

	return count
}

func BruteForceWithImprovedDP(A []int, L, R int) int {
	maxs := make([]int, len(A))
	count := 0
	max := 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			if i == j {
				max = A[j]
			} else {
				if A[j] > maxs[j-1] {
					max = A[j]
				} else {
					max = maxs[j-1]
				}
			}
			if L <= max && max <= R {
				count++
			}
			maxs[j] = max
		}
	}

	return count
}

// won't AC, memory limit exceeded
func BruteForceWithDP(A []int, L, R int) int {
	var max [][]int
	for i := 0; i < len(A); i++ {
		max = append(max, make([]int, len(A)-i))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; i+j < len(A); j++ {
			if j == 0 {
				max[i][0] = A[i]
			} else {
				if A[i+j] > max[i][j-1] {
					max[i][j] = A[i+j]
				} else {
					max[i][j] = max[i][j-1]
				}
			}
		}
	}

	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; i+j < len(A); j++ {
			if L <= max[i][j] && max[i][j] <= R {
				count++
			}
		}
	}

	return count
}

// won't AC, time limit exceeded
func BruteForce(A []int, L, R int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			max := 0
			for x := i; x <= j; x++ {
				if A[x] > max {
					max = A[x]
				}
			}
			if L <= max && max <= R {
				count++
			}
		}
	}

	return count
}

// @lc code=end


import (
	"sort"
)

/*
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (23.94%)
 * Total Accepted:    307.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return naive(nums1, nums2)
}

// merge two slice, sort, calculate median
// time complexity: (m+n)log(m+n)
func naive(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	length := len(merged)

	sort.Sort(sort.IntSlice(merged))

	if length&1 == 1 {
		return float64(merged[(length-1)/2])
	} else {
		return (float64(merged[length/2]) + float64(merged[length/2-1])) / 2
	}
}

// 1. If (length(A)≤2 or length(B)≤2) or (A.last≤B.first or B.last≤A.first) calculate median and return.
// 2. Set Am=median(A),Bm=median(B) and compare them. If Am=Bm return result.
// 		If Am<Bm then discard first half of A and the same amount of elements from the second half of B.
// 		else if Am>Bm then discard second half of A and the same amount of elements from the first half of B.
// 3. Goto 1
func formal(nums1 []int, nums2 []int) float64 {
	for {
		l1 := len(nums1)
		l2 := len(nums2)

		if l1 <= 2 || l2 <= 2 || nums1[l1-1] <= nums2[0] || nums2[l2-1] <= nums1[0] {
			break
		}

		m1 := median(nums1)
		m2 := median(nums2)
		count := min(l1/2, l2/2)
		if m1 > m2 {
			nums1 = nums1[0 : l1-count]
			nums2 = nums2[count:]
		} else if m1 < m2 {
			nums1 = nums1[count:]
			nums2 = nums2[0 : l2-count]
		} else {
			return m1
		}
	}

	return naive(nums1, nums2)
}

func median(slice []int) float64 {
	length := len(slice)
	if length&1 == 1 {
		return float64(slice[length-1] / 2)
	} else {
		return (float64(slice[length/2]) + float64(slice[length/2-1])) / 2
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

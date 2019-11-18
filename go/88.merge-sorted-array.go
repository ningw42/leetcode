/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	for i < m + j && j < n {
		fmt.Println(i, j)
		fmt.Println(nums1)
		fmt.Println(nums2)
		if nums1[i] >= nums2[j] {
			for k := m + j; k > i; k-- {
				nums1[k] = nums1[k-1]
			}
			nums1[i] = nums2[j]
			i++
			j++
		} else {
			i++
		}
	}
	for j < n {
		nums1[i] = nums2[j]
		i++
		j++
	}
}
// @lc code=end


/*
 * [461] Hamming Distance
 *
 * https://leetcode.com/problems/hamming-distance/description/
 *
 * algorithms
 * Easy (69.46%)
 * Total Accepted:    190.8K
 * Total Submissions: 274.8K
 * Testcase Example:  '1\n4'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Given two integers x and y, calculate the Hamming distance.
 *
 * Note:
 * 0 ≤ x, y < 231.
 *
 *
 * Example:
 *
 * Input: x = 1, y = 4
 *
 * Output: 2
 *
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 * ⁠      ↑   ↑
 *
 * The above arrows point to positions where the corresponding bits are
 * different.
 *
 *
 */
func hammingDistance(x int, y int) int {
	diff := x ^ y
	count := 0

	for diff != 0 {
		if diff&1 == 1 {
			count += 1
		}
		diff = diff >> 1
	}

	return count
}

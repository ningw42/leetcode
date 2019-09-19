/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
func coinChange(coins []int, amount int) int {
	r := []int{0}
	for i := 0; i < amount; i++ {
		r = append(r, -1)
	}
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				count := r[i-coin]
				if count != -1 && count < min {
					min = count + 1
				}
			}
		}
		if min != math.MaxInt32 {
			r[i] = min
		}
	}

	return r[amount]
}


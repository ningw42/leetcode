/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	seq := []byte{'1'}
	for i := 1; i < n; i++ {
		seq = say(seq)
	}

	return string(seq)
}

func say(seq []byte) []byte {
	var ret []byte
	var count int
	var char byte

	char = seq[0]
	count = 1
	for i := 1; i < len(seq); i++ {
		if seq[i] != char {
			ret = append(ret, byte(count+'0'), char)
			char = seq[i]
			count = 1
		} else {
			count++
		}
	}
	ret = append(ret, byte(count+'0'), char)

	return ret
}

// @lc code=end


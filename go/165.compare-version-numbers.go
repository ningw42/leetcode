/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	// pad '0' to make the length of v1 and v2 identical
	var length, diff int
	if len(v1) > len(v2) {
		length = len(v1)
		diff = len(v1) - len(v2)
		for i := 0; i < diff; i++ {
			v2 = append(v2, "0")
		}
	} else if len(v1) < len(v2) {
		length = len(v2)
		diff = len(v2) - len(v1)
		for i := 0; i < diff; i++ {
			v1 = append(v1, "0")
		}
	} else {
		length = len(v1)
	}

	for i := 0; i < length; i++ {
		vs1, _ := strconv.Atoi(v1[i])
		vs2, _ := strconv.Atoi(v2[i])
		if vs1 > vs2 {
			return 1
		}
		if vs1 < vs2 {
			return -1
		}
	}
	return 0
}

// @lc code=end


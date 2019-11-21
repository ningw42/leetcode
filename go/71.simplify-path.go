/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	segments := strings.Split(path, "/")

	var canonical []string
	for _, segment := range segments {
		if segment != "" && segment != "." {
			if segment == ".." {
				if len(canonical) != 0 {
					canonical = canonical[0:len(canonical)-1]
				}
			} else {
				canonical = append(canonical, segment)
			}
		}
	}

	return "/" + strings.Join(canonical, "/")
}
// @lc code=end


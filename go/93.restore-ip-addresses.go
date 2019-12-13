/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
var ips []string
func restoreIpAddresses(s string) []string {
	ips = nil
	restore([]byte(s), []string{})
	return ips
}
func restore(s []byte, segments []string) {
	if len(s) == 0 && len(segments) == 4 {
		ips = append(ips, strings.Join(segments, "."))
	} else {
		if len(s) == 0 || len(segments) == 4 {
			return
		}
		restore(s[1:], append(segments, string(s[0:1])))
		if s[0] == '0' || len(s) < 2 {
			return
		}
		restore(s[2:], append(segments, string(s[0:2])))
		if segment, _ := strconv.Atoi(string(s[0:3])); segment > 255 || len(s) < 3 {
			return
		}
		restore(s[3:], append(segments, string(s[0:3])))
	}
}
// @lc code=end


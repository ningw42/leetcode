/*
 * @lc app=leetcode id=937 lang=golang
 *
 * [937] Reorder Data in Log Files
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	var slogs [][]string
	var dlogs []string
	for _, log := range logs {
		parsed := parse(log)
		if parsed[1][0] >= '0' && parsed[1][0] <= '9' {
			dlogs = append(dlogs, log)
		} else {
			slogs = append(slogs, parsed)
		}
	}

	sort.Slice(slogs, func(i, j int) bool {
		result := strings.Compare(slogs[i][1], slogs[j][1])
		if result < 0 {
			return true
		} else if result > 0 {
			return false
		} else {
			return strings.Compare(slogs[i][0], slogs[j][0]) < 0
		}
	})

	ret := make([]string, len(logs))
	var i int
	for _, slog := range slogs {
		ret[i] = strings.Join(slog, " ")
		i++
	}
	for _, dlog := range dlogs {
		ret[i] = dlog
		i++
	}

	return ret
}

func parse(s string) []string {
	space := strings.IndexByte(s, ' ')
	return []string{s[:space], s[space+1:]}
}
// @lc code=end


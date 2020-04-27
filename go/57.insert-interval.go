func insert(intervals [][]int, newInterval []int) [][]int {
	return merge(append(intervals, newInterval))
}

type Intervals [][]int
func (i Intervals) Len() int {
	return len(i)
}
func (i Intervals) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}
func (i Intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort by left boundry asceding
	sort.Sort(Intervals(intervals))
	for i := 0; i < len(intervals)-1; {
		// if intersects
		if intervals[i][1] >= intervals[i+1][0] {
			intervals = append(
				intervals[:i], 
				append(
					[][]int{[]int{intervals[i][0], maxInt(intervals[i][1], intervals[i+1][1])}},
					intervals[i+2:]...
				)...
			)
		} else {
			i++
		}
	}

	return intervals
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	indexes, sources, targets = reorder(indexes, sources, targets)
	var index, shift, shiftIndex int
	var source, target string
	shift = 0
	str := S
	for len(indexes) != 0 {
		index = indexes[0]
		shiftIndex = indexes[0] + shift
		source = sources[0]
		target = targets[0]

		if source == S[index:index+len(source)] {
			str = str[0:shiftIndex] + target + str[shiftIndex+len(source):]
			shift += len(target) - len(source)
		}

		indexes = indexes[1:]
		sources = sources[1:]
		targets = targets[1:]
	}

	return str
}

type Step struct {
	Index  int
	Source string
	Target string
}

type Steps []Step

func (s Steps) Len() int {
	return len(s)
}
func (s Steps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Steps) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

func reorder(indexes []int, sources, targets []string) ([]int, []string, []string) {
	var steps []Step
	for i := 0; i < len(indexes); i++ {
		steps = append(steps, Step{
			Index:  indexes[i],
			Source: sources[i],
			Target: targets[i],
		})
	}

	sort.Sort(Steps(steps))

	var nIndexes []int
	var nSources, nTargets []string
	for i := 0; i < len(steps); i++ {
		nIndexes = append(nIndexes, steps[i].Index)
		nSources = append(nSources, steps[i].Source)
		nTargets = append(nTargets, steps[i].Target)
	}

	return nIndexes, nSources, nTargets
}

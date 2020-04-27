
var count map[string]int

type WordHeap []string

func (wh WordHeap) Len() int {
	return len(wh)
}
func (wh WordHeap) Less(i, j int) bool {
	return !MoreFrequent(wh[i], wh[j])
}
func (wh WordHeap) Swap(i, j int) {
	wh[i], wh[j] = wh[j], wh[i]
}
func (wh *WordHeap) Push(x interface{}) {
	*wh = append(*wh, x.(string))
}
func (wh *WordHeap) Pop() interface{} {
	old := *wh
	n := len(old)
	x := old[n-1]
	*wh = old[0 : n-1]
	return x
}

func MoreFrequent(a, b string) bool {
	aCount := count[a]
	bCount := count[b]
	if aCount == bCount {
		return a <= b
	} else if aCount > bCount {
		return true
	} else {
		return false
	}
}

// 1. build min heap with size k
// 2. scan all words with that min heap
// 3. pop all heap element in reversed order
func topKFrequent(words []string, k int) []string {
	count = make(map[string]int)
	h := &WordHeap{}
	i := 0
	for len(count) < k && i < len(words) {
		if _, exists := count[words[i]]; exists {
			count[words[i]]++
		} else {
			h.Push(words[i])
			count[words[i]] = 1
		}
		i++
	}
	heap.Init(h)

	for i < len(words) {
		if _, exists := count[words[i]]; exists {
			count[words[i]]++
		} else {
			count[words[i]] = 1
		}
		if MoreFrequent(words[i], (*h)[0]) {
			// fmt.Printf("%s is more frequent than %s\n", words[i], (*h)[0])
			exists := false
			for index, word := range *h {
				if word == words[i] {
					heap.Fix(h, index)
					exists = true
					break
				}
			}
			if !exists {
				heap.Pop(h)
				heap.Push(h, words[i])
			}
		}
		i++
	}

	result := make([]string, h.Len())
	for i := h.Len() - 1; h.Len() > 0; i-- {
		result[i] = heap.Pop(h).(string)
	}

	return result
}

import "container/heap"

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start

var count map[byte]int

type ByteHeap []byte

func (h ByteHeap) Len() int {
	return len(h)
}
func (h ByteHeap) Less(i, j int) bool {
	bi := h[i]
	bj := h[j]
	ci := count[bi]
	cj := count[bj]

	if ci < cj {
		return false
	} else if ci == cj {
		return bi < bj
	} else {
		return true
	}
}
func (h ByteHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ByteHeap) Pop() interface{} {
	length := h.Len()
	result := (*h)[length-1]
	*h = (*h)[0 : length-1]
	return result
}
func (h *ByteHeap) Push(x interface{}) {
	*h = append(*h, x.(byte))
}

func frequencySort(s string) string {
	h := &ByteHeap{}
	count = make(map[byte]int)
	for _, r := range s {
		c := byte(r)
		if _, exists := count[c]; exists {
			count[c]++
		} else {
			count[c] = 1
		}
		h.Push(c)
	}
	heap.Init(h)

	var result []byte
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(byte))
	}

	return string(result)
}

// @lc code=end


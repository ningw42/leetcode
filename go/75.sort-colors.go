import (
	"container/heap"
)

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start

// 1. two pass, count all [0, 1, 2], and overide with order
// 2. heap sort

// An IntHeap is a max-heap of ints.
type IntHeap struct {
	d []int
}

func (h *IntHeap) Len() int           { return len(h.d) }
func (h *IntHeap) Less(i, j int) bool { return h.d[i] >= h.d[j] }
func (h *IntHeap) Swap(i, j int)      { h.d[i], h.d[j] = h.d[j], h.d[i] }
func (h *IntHeap) Push(x interface{}) {
	h.d = append(h.d, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := h.d
	n := len(old)
	x := old[n-1]
	h.d = old[0 : n-1]
	return x
}

func sortColors(nums []int) {
	h := &IntHeap{d: nums}
	heap.Init(h)
	for i := len(nums) - 1; i >= 0; i-- {
		max := heap.Pop(h).(int)
		nums[i] = max
	}
}

// @lc code=end


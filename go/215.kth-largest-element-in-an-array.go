/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
import (
	"container/heap"
)

// copied from https://golang.org/pkg/container/heap/#example__intHeap
// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] >= h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// get the k-th largest item by popping from a max heap k times
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	for _, num := range nums {
		h.Push(num)
	}
	heap.Init(h)

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}
	result := heap.Pop(h).(int)

	return result
}

// @lc code=end


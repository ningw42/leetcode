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

func findKthLargest(nums []int, k int) int {
	// min heap nlog(k)
	// max heap klog(n)
	// quick select n on averege and n^2 in worst case
	return quickSelect(nums, 0, len(nums) - 1, k)
}

// get the k-th largest item by popping from a max heap k times
func maxHeap(nums []int, k int) int {
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

func quickSelect(nums []int, left, right, k int) int {
    if left == right {
        return nums[left]
    }
    pivotIndex := partition(nums, left, right)
    if pivotIndex > k {
        return quickSelect(nums, left, pivotIndex - 1, k)
    } else if pivotIndex < k {
        return quickSelect(nums, pivotIndex + 1, right, k)
    } else {
        return nums[k]
    }
}

func partition(nums []int, left, right int) int {
    pivotIndex := medianOfThree(nums, left, right)
    pivotValue := nums[pivotIndex]
    // swap nums[pivotIndex] and nums[right]
    swap(nums, pivotIndex, right)
    // select num that > pivotValue
    currentIndex := left
    for i := left; i < right; i++ {
        if nums[i] > pivotValue {
            // swap nums[i] and nums[currentIndex]
            swap(nums, currentIndex, i)
            currentIndex++
        }
    }
    // restore pivot's location
    swap(nums, currentIndex, right)
    return currentIndex
}

func medianOfThree(nums []int, left, right int) int {
    mid := (left + right) / 2
    // sort nums[left], nums[mid], nums[right]
    if nums[left] < nums[right] {
        swap(nums, left, right)
    }
    if nums[mid] < nums[right] {
        swap(nums, mid, right)
    }
    if nums[left] < nums[mid] {
        swap(nums, left, mid)
    }
    fmt.Println(nums[left], nums[mid], nums[right])
    return mid
}

func swap(nums []int, i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end


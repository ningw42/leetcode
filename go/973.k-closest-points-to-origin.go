/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

// @lc code=start
// max heap
type PriorityQueue [][]int
func (pq PriorityQueue) Len() int {
    return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
    return distance(pq[i]) > distance(pq[j])
}
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.([]int))
}
func (pq *PriorityQueue) Pop() interface{} {
    length := pq.Len()
    element := (*pq)[length-1]
    *pq = (*pq)[:length-1]
    return element
}

func kClosest(points [][]int, K int) [][]int {
    var mh PriorityQueue
    heapContainer := make([][]int, K)
    for i, point := range points {
        if i < K {
            heapContainer[i] = point
            continue
        }
        if i == K {
            mh = PriorityQueue(heapContainer)
            heap.Init(&mh)
        }
        if distance(point) < distance(heapContainer[0]) {
            heap.Pop(&mh)
            heap.Push(&mh, point)
        }
    }
 
    return heapContainer
}

func distance(p []int) int {
    return p[0]*p[0] + p[1]*p[1]
}
// @lc code=end


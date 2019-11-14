import (
	"container/heap"
)

/*
 * @lc app=leetcode id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start

type Project struct {
	Profit  int
	Capital int
}

type MinCaptitalHeap []Project
type MaxProfitHeap []Project

func (p MinCaptitalHeap) Len() int {
	return len(p)
}
func (p MinCaptitalHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p MinCaptitalHeap) Less(i, j int) bool {
	return p[i].Capital < p[j].Capital
}
func (p *MinCaptitalHeap) Push(x interface{}) {
	*p = append(*p, x.(Project))
}
func (p *MinCaptitalHeap) Pop() interface{} {
	length := p.Len()
	x := (*p)[length-1]
	*p = (*p)[0 : length-1]
	return x
}

func (p MaxProfitHeap) Len() int {
	return len(p)
}
func (p MaxProfitHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p MaxProfitHeap) Less(i, j int) bool {
	return p[i].Profit >= p[j].Profit
}
func (p *MaxProfitHeap) Push(x interface{}) {
	*p = append(*p, x.(Project))
}
func (p *MaxProfitHeap) Pop() interface{} {
	length := p.Len()
	x := (*p)[length-1]
	*p = (*p)[0 : length-1]
	return x
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	minCaptitalHeap := &MinCaptitalHeap{}
	maxProfitHeap := &MaxProfitHeap{}
	for i, p := range Profits {
		*minCaptitalHeap = append(*minCaptitalHeap, Project{
			Profit:  p,
			Capital: Capital[i],
		})
	}

	heap.Init(minCaptitalHeap)
	heap.Init(maxProfitHeap)

	for k > 0 {
		for minCaptitalHeap.Len() > 0 {
			if (*minCaptitalHeap)[0].Capital <= W {
				heap.Push(maxProfitHeap, heap.Pop(minCaptitalHeap))
			} else {
				break
			}
		}
		if maxProfitHeap.Len() > 0 {
			project := heap.Pop(maxProfitHeap).(Project)
			W += project.Profit
			k--
		} else {
			break
		}
	}

	return W
}

// @lc code=end


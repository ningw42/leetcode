/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(begin, end int) []*TreeNode {
	if begin == end {
		return []*TreeNode{&TreeNode{Val: begin}}
	}
	var ret []*TreeNode
	for i := begin; i <= end; i++ {
		var lefts, rights []*TreeNode
		if i == begin {
			lefts = []*TreeNode{nil}
			rights = generate(i+1, end)
		} else if i == end {
			lefts = generate(begin, i-1)
			rights = []*TreeNode{nil}
		} else {
			lefts = generate(begin, i-1)
			rights = generate(i+1, end)
		}
		for _, left := range lefts {
			for _, right := range rights {
				ret = append(ret, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return ret
}

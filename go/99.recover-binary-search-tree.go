/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
    if root == nil {
        return
    }
    raw := traversal(root)
	sorted := make([]int, len(raw))
	copy(sorted, raw)
	sort.Ints(sorted)

	var a, b *TreeNode
	for i, e := range raw {
		if e != sorted[i] {
			a = find(root, e)
			b = find(root, sorted[i])
			break
		}
	}

	t := a.Val
	a.Val = b.Val
	b.Val = t
}

func traversal(t *TreeNode) []int {
	if t == nil {
        return []int{}
	} else {
        return append(append(traversal(t.Left), t.Val), traversal(t.Right)...)
	}
}

func find(t *TreeNode, e int) *TreeNode {
	if t == nil {
		return nil
	} else {
		if t.Val == e {
			return t
		} else {
			n := find(t.Left, e)
			if n == nil {
				return find(t.Right, e)
			} else {
				return n
			}
		}
	}
}
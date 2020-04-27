/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(N int) []*TreeNode {
    if N % 2 == 0 {
        return nil
    }
    return find(N, map[int][]*TreeNode{})
}

func find(n int, m map[int][]*TreeNode) []*TreeNode {
    if n == 1 {
        m[1] = []*TreeNode{&TreeNode{}}
        return []*TreeNode{&TreeNode{}}
    }
    var result []*TreeNode
    for left := 1; left + 1 <= n - 1; left = left + 2 {
        ls := find(left, m)
        rs := find(n-1-left, m)
        for _, l := range ls {
            for _, r := range rs {
                result = append(result, &TreeNode{
                    Left: l,
                    Right: r,
                })
            }
        }
    }
    m[n] = result
    return result
}
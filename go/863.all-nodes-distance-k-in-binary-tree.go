/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    // build graph
    graph := make([][]bool, 501)
    for i := 0; i <= 500; i++ {
        graph[i] = make([]bool, 501)
    }
    traversal(root, graph)
    
    // bfs
    queue := []int{target.Val}
    nextQueue := []int{}
    distance := 0;
    for distance < K && len(queue) > 0 {
        for len(queue) > 0 {
            from := queue[0]
            for to, edge := range graph[from] {
                if edge {
                    nextQueue = append(nextQueue, to)
                    graph[from][to] = false
                    graph[to][from] = false
                }
            }
            queue = queue[1:]
        }
        distance++
        queue, nextQueue = nextQueue, queue
    }
    
    return queue
}

func traversal(root *TreeNode, graph [][]bool) {
    if root.Left != nil {
        graph[root.Val][root.Left.Val] = true
        graph[root.Left.Val][root.Val] = true
        traversal(root.Left, graph)
    }
    if root.Right != nil {
        graph[root.Val][root.Right.Val] = true
        graph[root.Right.Val][root.Val] = true
        traversal(root.Right, graph)
    }
}
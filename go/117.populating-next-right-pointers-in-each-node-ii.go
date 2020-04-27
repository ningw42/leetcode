/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    doConnect(root, 0, map[int]*Node{})
    return root
}

func doConnect(root *Node, level int, nexts map[int]*Node) {
    if root == nil {
        return
    } else {
        // link next
        root.Next = nexts[level]
        // traverse right subtree and update 'nexts'
        doConnect(root.Right, level+1, nexts)
        // traverse left subtree and update 'nexts' 
        doConnect(root.Left, level+1, nexts)
        // override 'nexts' on current level
        nexts[level] = root
    }
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    node, parent := findNode(root, nil, key)
    if node == nil {
        // no node with given key
        return root
    } else {
        // node with given key is found
        // try to delete node
        if node == root {
            // delete root
            // find min in right subtree of the found node to replace node
            min, _ := findMin(node.Right, node)
            if min == nil {
                // cannot find min
                // that is, node's right subtree is empty
                // since we are deleting root, just return root's left subtree will do
                return root.Left
            } else {
                // left subtree of original root stays unchanged
                // delete min from the right subtree of original root
                // min will be the new root
                min.Left, min.Right = root.Left, deleteNode(node.Right, min.Val)
                return min
            }
        } else {
            // delete non-root
            min, _ := findMin(node.Right, node)
            if min == nil {
                // right subtree of the targeting node is empty
                // connect node's left subtree to node's parent
                if node == parent.Left {
                    parent.Left = node.Left
                } else {
                    parent.Right = node.Left
                }
                return root
            } else {
                // delete min from node's right subtree
                // connect parent and the min
                if node == parent.Left {
                    parent.Left, min.Left, min.Right = min, node.Left, deleteNode(node.Right, min.Val)
                } else {
                    parent.Right, min.Left, min.Right = min, node.Left, deleteNode(node.Right, min.Val)
                }
                return root
            }
        }
    }
}

// find node with given key in subtree
func findNode(current, parent *TreeNode, key int) (*TreeNode, *TreeNode) {
    if current == nil {
        return nil, parent
    } else {
        if current.Val > key {
            return findNode(current.Left, current, key)
        } else if current.Val < key {
            return findNode(current.Right, current, key)
        } else {
            return current, parent
        }
    }
}

// find node with min value in subtree
func findMin(current, parent *TreeNode) (*TreeNode, *TreeNode) {
    if current == nil {
        return nil, parent
    } else {
        if current.Left != nil {
            return findMin(current.Left, current)
        } else {
            return current, parent
        }
    }
}
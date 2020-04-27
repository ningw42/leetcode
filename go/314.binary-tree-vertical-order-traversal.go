/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m map[int][][]int
var leftMostColumn int

func verticalOrder(root *TreeNode) [][]int {
    m = make(map[int][][]int)
    leftMostColumn = 0
    
    traverse(root, 0, 0)
    
    ret := make([][]int, len(m))
    for column, values := range m {
        // sort by depth ascending
        // a stable sort is required to preserve the left-to-right order inside same column
        sort.SliceStable(values, func(i, j int) bool {
            return values[i][1] < values[j][1]
        })
        var columnValues []int
        for _, value := range values {
            columnValues = append(columnValues, value[0])
        }
        ret[column-leftMostColumn] = columnValues
    }
    
    return ret
}

func traverse(root *TreeNode, column int, depth int) {
    if root == nil {
        return 
    } else {
        leftMostColumn = min(leftMostColumn, column)
        // value-depth pairs
        m[column] = append(m[column], []int{root.Val, depth})
        traverse(root.Left, column-1, depth+1)
        traverse(root.Right, column+1, depth+1)
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
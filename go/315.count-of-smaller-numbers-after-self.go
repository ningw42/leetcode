func countSmaller(nums []int) []int {
    return buildBST(nums)
}

// building a BST while iterating from nums[len(nums)-1] to nums[0]
// count number of existing nodes with smaller value than nums[i]
func buildBST(nums []int) []int {
    if len(nums) == 0 {return nil}

    counts := make([]int, len(nums))
    root := &MyTreeNode{
        value: nums[len(nums)-1],
        count: 1,
    }
    for i := len(nums) - 2; i >= 0; i-- {
        counts[i] = appendNode(root, &MyTreeNode{
            value: nums[i],
            count: 1,
        })
    }
    return counts
}

type MyTreeNode struct {
    value int
    count int
    left *MyTreeNode
    right *MyTreeNode
}

func appendNode(root, node *MyTreeNode) int {
    var parent *MyTreeNode
    var leftSubTree bool
    var count int
    for root != nil {
        if root.value < node.value {
            count += root.count
            root, parent = root.right, root
            leftSubTree = false
        } else {
            // root.value >= node.value
            root.count++
            root, parent = root.left, root
            leftSubTree = true
        }
    }
    if leftSubTree {
        parent.left = node
    } else {
        parent.right = node
    }
    return count
}

// count element on the right of nums[i] moved to nums[i] left in a stable sort
// any stable sort will do
func mergeSort(nums []int) []int {
    // TODO
    return nil
}

// O(N^2), AC but slow
func bruteforce(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }
    
    counts := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] > nums[j] {
                counts[i]++
            }
        }
    }
    
    return counts;
}

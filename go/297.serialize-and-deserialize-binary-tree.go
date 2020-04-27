/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return strings.Join(this.traverse(root), ",")
}

// prefix
func (c *Codec) traverse(root *TreeNode) []string {
    if root == nil {
        return []string{"n"}
    } else {
        return append(
            []string{strconv.Itoa(root.Val)},
            append(c.traverse(root.Left), c.traverse(root.Right)...)...,
        )
    }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    // fmt.Println(data)
    values := strings.Split(data, ",")
    root, _ := this.build(values)
    return root
}

func (c *Codec) build(values []string) (*TreeNode, int) {
    if values[0] == "n" {
        return nil, 1
    } else {
        value, _ := strconv.Atoi(values[0])
        node := &TreeNode{Val: value}
        left, leftCount := c.build(values[1:])
        right, rightCount := c.build(values[1+leftCount:])
        node.Left = left
        node.Right = right
        return node, 1 + leftCount + rightCount
    }
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
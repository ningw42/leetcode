/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    prefix []string
    infix []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    this.prefixTraverse(root)
    this.infixTraverse(root)
    return strings.Join(this.prefix, ",") + ";" + strings.Join(this.infix, ",")
}

func (c *Codec) prefixTraverse(root *TreeNode) {
    if root == nil {
        return
    } else {
        c.prefix = append(c.prefix, strconv.Itoa(root.Val))
        c.prefixTraverse(root.Left)
        c.prefixTraverse(root.Right)
    }
}

func (c *Codec) infixTraverse(root *TreeNode) {
    if root == nil {
        return
    } else {
        c.infixTraverse(root.Left)
        c.infix = append(c.infix, strconv.Itoa(root.Val))
        c.infixTraverse(root.Right)
    }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    tmp := strings.Split(data, ";")
    // fmt.Println(tmp[0])
    // fmt.Println(tmp[1])
    // prefix and infix may be [""], filtered by parse()
    prefix := parse(strings.Split(tmp[0], ","))
    infix := parse(strings.Split(tmp[1], ","))
    return this.build(prefix, infix)
}

func (c *Codec) build(prefix, infix []int) *TreeNode {
    if len(prefix) == 0 {
        return nil
    } else {
        var i int
        for infix[i] != prefix[0] {
            i++
        }
        return &TreeNode{
            Val: prefix[0],
            Left: c.build(prefix[1:i+1], infix[:i]),
            Right: c.build(prefix[i+1:], infix[i+1:]),
        }
    }
}

func parse(ss []string) []int {
    var is []int
    for _, s := range ss {
        i, err := strconv.Atoi(s)
        if err != nil {
            return nil
        }
        is = append(is, i)
    }
    return is
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
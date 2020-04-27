/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Codec struct {
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
    if root == nil {
        return ""
    } else {
        if len(root.Children) == 0 {
            return fmt.Sprintf("%d[]", root.Val)
        } else {
            var builder strings.Builder
            builder.WriteString(strconv.Itoa(root.Val))
            builder.WriteByte('[')
            for _, child := range root.Children {
                builder.WriteString(this.serialize(child))
            }
            builder.WriteByte(']')
            return builder.String()
        }
    }
}

func (this *Codec) deserialize(data string) *Node {
    // fmt.Println(data)
    if data == "" {
        return nil
    } else {
        // parse root value
        var i int
        for i = 0; i < len(data) && isDigit(data[i]); i++ {}
        value, _ := strconv.Atoi(data[:i])
        node := &Node{Val: value}
        
        // parse children
        i++
        if i == len(data) - 1 {
            // no child
            return node
        } else {
            for i < len(data) - 1 {
                // each child
                started := false
                count := 0
                j := i
                for {
                    if data[j] == '[' {
                        count++
                        if started == false {
                            started = true
                        }
                    } else if data[j] == ']' {
                        count--
                    }
                    j++
                    if count == 0 && started {
                        node.Children = append(node.Children, this.deserialize(data[i:j]))
                        break
                    }
                }
                i = j
            }
            return node
        }
    }
}

func isDigit(c byte) bool {
    return '0' <= c && c <= '9'
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
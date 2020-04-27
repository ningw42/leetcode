type MyLinkedList struct {
    head *MyListNode
    tail *MyListNode
}

type MyListNode struct {
    Value int
    Prev *MyListNode
    Next *MyListNode
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    pos := this.get(index)
    if pos == nil {
        return -1
    } else {
        return pos.Value
    }
}

func (this *MyLinkedList) get(index int) *MyListNode {
    curr := this.head
    for count := 0; curr != nil && count < index; count++ {
        curr = curr.Next
    }
    return curr
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    if this.head == nil {
        node := &MyListNode{Value:val}
        this.head = node
        this.tail = node
    } else {
        node := &MyListNode{
            Value: val,
            Next: this.head,
        }
        this.head, this.head.Prev = node, node
    }
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    if this.tail == nil {
        node := &MyListNode{Value:val}
        this.head = node
        this.tail = node
    } else {
        node := &MyListNode{
            Value: val,
            Prev: this.tail,
        }
        this.tail, this.tail.Next = node, node
    }
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    pos := this.get(index)
    if pos == this.head {
        this.AddAtHead(val)
    } else if pos == nil {
        this.AddAtTail(val)
    } else {
        node := &MyListNode{
            Value: val,
            Prev: pos.Prev,
            Next: pos,
        }
        pos.Prev, pos.Prev.Next = node, node
    }
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    pos := this.get(index)
    if pos == nil {
        return
    }
    if pos == this.head && pos == this.tail {
        this.head = nil
        this.tail = nil
    } else if pos == this.head {
        this.head, this.head.Next.Prev = this.head.Next, nil
    } else if pos == this.tail {
        this.tail, this.tail.Prev.Next = this.tail.Prev, nil
    } else {
        pos.Prev.Next, pos.Next.Prev = pos.Next, pos.Prev
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
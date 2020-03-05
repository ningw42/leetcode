/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
 type Node struct {
    Key   int
    Value int
    Prev  *Node
    Next  *Node
}

type LRUCache struct {
    m map[int]*Node
    head *Node
    tail *Node
    capacity int
}


func Constructor(capacity int) LRUCache {
    h := &Node{}
    t := &Node{}
    h.Next = t
    t.Prev = h
    return LRUCache{
        m: make(map[int]*Node),
        head: h,
        tail: t,
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if node, exists := this.m[key]; exists {
        this.moveToHead(node)
        return node.Value
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if node, exists := this.m[key]; exists {
        node.Value = value
        this.moveToHead(node)
    } else {
        // append to head
        node := &Node{Key: key, Value: value}
        this.m[key] = node
        this.head.Next, node.Next, this.head.Next.Prev, node.Prev = node, this.head.Next, node, this.head
        if len(this.m) > this.capacity {
            delete(this.m, this.tail.Prev.Key)
            this.tail.Prev, this.tail.Prev.Prev.Next = this.tail.Prev.Prev, this.tail
        }
    }
}

func (lru *LRUCache) debug() {
    var vals []int
    for head := lru.head; head != nil; head = head.Next {
        vals = append(vals, head.Value)
    }
    fmt.Printf("%v\n", vals)
    fmt.Println("---------")
}

func (this *LRUCache) moveToHead(node *Node) {
        node.Prev.Next, node.Next.Prev = node.Next, node.Prev
        node.Prev, node.Next, this.head.Next, this.head.Next.Prev = this.head, this.head.Next, node, node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


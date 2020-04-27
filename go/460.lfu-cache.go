// doubly linked list node with key, value and its access count
type Node struct {
    key   int
    value int
    count int
    prev  *Node
    next  *Node
}

func newNode() *Node {
    return &Node{
        prev: nil,
        next: nil,
    }
}

// doubly linked list with dummy head and tail
type DoublyLinkedList struct {
    head *Node
    tail *Node
}
func newDoublyLinkedList() *DoublyLinkedList {
    head := newNode()
    tail := newNode()
    head.next, tail.prev = tail, head
    return &DoublyLinkedList{
        head: head,
        tail: tail,
    }
}
// tell if current doubly linked list is empty
func (list *DoublyLinkedList) empty() bool {
    return list.head.next == list.tail && list.tail.prev == list.head
}
// append node to the head of doubly linked list
func (list *DoublyLinkedList) prepend(node *Node) {
    list.head.next, list.head.next.prev, node.prev, node.next = node, node, list.head, list.head.next
}
// pop node from the tail of doubly linked list
func (list *DoublyLinkedList) pop() *Node {
    if list.empty() {
        panic("pop on empty list")
    } else {
        node := list.tail.prev
        list.tail.prev, list.tail.prev.prev.next = list.tail.prev.prev, list.tail
        return node
    }
}

// count-leveled doubly linked list with hash map
// both Get and Put are O(1)
type LFUCache struct {
    capacity int // total capacity
    minCount int // current min access count
    m map[int]*Node // 'key' as key, Node as value
    c map[int]*DoublyLinkedList // access count as key, doubly linked list as value, all nodes in list share the same access count
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity: capacity,
        minCount: 0,
        m: make(map[int]*Node),
        c: make(map[int]*DoublyLinkedList),
    }
}


func (this *LFUCache) Get(key int) int {
    if this.capacity == 0 {
        return -1
    }
    if node, exists := this.m[key]; !exists {
        // not exists
        return -1
    } else {
        // exists
        value := node.value
        oldCount := node.count
        // detach node from list with its old count
        node.prev.next, node.next.prev = node.next, node.prev
        if oldCount == this.minCount && this.c[this.minCount].empty() {
            this.minCount++
        }
        // advance node's count value
        node.count++
        // attach node to the beginning of the list with its new count
        if _, listExists := this.c[node.count]; !listExists {
            this.c[node.count] = newDoublyLinkedList()
        }
        this.c[node.count].prepend(node)
        return value
    }
}


func (this *LFUCache) Put(key int, value int)  {
    if this.capacity == 0 {
        return
    }
    if node, exists := this.m[key]; !exists {
        // not exists
        node := &Node{
            key: key,
            value: value,
            count: 1,
            prev: nil,
            next: nil,
        }
        // eviction
        if len(this.m) == this.capacity {
            // exceed capacity, evict one entry
            evictedNode := this.c[this.minCount].pop()
            delete(this.m, evictedNode.key)
        }
        // put node to hashmap
        this.m[key] = node
        if _, listExists := this.c[1]; !listExists {
            this.c[1] = newDoublyLinkedList()
        }
        // put node to count indexed doubly linked list
        this.c[1].prepend(node)
        this.minCount = 1
    } else {
        // exists
        oldCount := node.count
        // update node value
        node.value = value
        // detach node from list with its old count
        node.prev.next, node.next.prev = node.next, node.prev
        if oldCount == this.minCount && this.c[oldCount].empty() {
            this.minCount++
        }
        // advance node count
        node.count++
        // attach node to the beginning of list with its new count
        if _, listExists := this.c[node.count]; !listExists {
            this.c[node.count] = newDoublyLinkedList()
        }
        this.c[node.count].prepend(node)
    }
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type Char struct {
    char byte
    count int
}

// max heap by Char.count
type CharHeap []*Char
func newCharHeap() *CharHeap {
    var container []*Char
    h := CharHeap(container)
    return &h
}
func (h CharHeap) Len() int {return len(h)}
func (h CharHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h CharHeap) Less(i, j int) bool {return h[i].count > h[j].count}
func (h *CharHeap) Push(x interface{}) {*h = append(*h, x.(*Char))}
func (h *CharHeap) Pop() interface{} {
    length := h.Len()
    x := (*h)[length-1]
    *h = (*h)[:length-1]
    return x
}

func reorganizeString(S string) string {
    // count for each char
    c := make(map[byte]int)
    for i := 0; i < len(S); i++ {
        c[S[i]]++
    }
  
    // put those couns into max heap
    h1 := newCharHeap()
    for char, count := range c {
        heap.Push(h1, &Char{
            char: char,
            count: count,
        })
    }
  
    // greedy: keep finding the char with largest count to append
    var bs []byte
    var stash []*Char
    for h1.Len() > 0 {
        if len(bs) == 0 {
            // the very first iteration
            char := heap.Pop(h1).(*Char)
            bs = append(bs, char.char)
            char.count--
            if char.count != 0 {
                heap.Push(h1, char)
            }
        } else {
            for h1.Len() > 0 {
                // pop a char with max count from heap
                char := heap.Pop(h1).(*Char)
                if char.char != bs[len(bs)-1] {
                    // if the char is not identical to the last one
                    bs = append(bs, char.char)
                    char.count--
                    if char.count != 0 {
                        // put back to heap if count is not zero
                        heap.Push(h1, char)
                    }
                    // put back all stashed chars during finding the current one
                    for _, stashedChar := range stash {
                        heap.Push(h1, stashedChar)
                    }
                    // reset stash chars
                    stash = nil
                    break
                } else {
                    // if the char is identical to the last one
                    // ignore and stash current one
                    stash = append(stash, char)
                }
            }
        }
    }

    if len(stash) > 0 {
        return ""
    } else {
        return string(bs)
    }
}
type Item struct {
    v string // value
    t int    // timestamp
}

type ItemSlice []Item

func (s *ItemSlice) Add(item Item) {
    s.AddAt(len(*s), item)
}

func (s *ItemSlice) AddAt(i int, item Item) {
    if i == len(*s) {
        *s = append(*s, item)
    } else {
        *s = append((*s)[:i], append([]Item{item}, (*s)[i:]...)...)
    }
}

func (s *ItemSlice) Get(i int) Item {
    return (*s)[i]
}

func (s *ItemSlice) Find(timestamp int) int {
    i := 0
    j := len(*s) - 1
    for i < j {
        k := i + (j-i) / 2 + 1
        if (*s)[k].t < timestamp {
            i = k
        } else if (*s)[k].t > timestamp {
            j = k - 1
        } else {
            return k
        }
    }
    if (*s)[i].t <= timestamp {
        return i
    } else {
        return -1
    }
}

type TimeMap struct {
    m map[string]*ItemSlice
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{
        m: make(map[string]*ItemSlice),
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if slice, exists := this.m[key]; exists {
        i := slice.Find(timestamp)
        this.m[key].AddAt(i+1, Item{v:value,t:timestamp})
    } else {
        newSlice := ItemSlice(make([]Item, 0))
        this.m[key] = &newSlice
        this.m[key].Add(Item{v:value,t:timestamp})
    }
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if slice, exists := this.m[key]; exists {
        i := slice.Find(timestamp)
        if i == -1 {
            return ""
        } else {
            return slice.Get(i).v
        }
    } else {
        return ""
    }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
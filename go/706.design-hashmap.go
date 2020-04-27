type MyHashMap struct {
    buckets [][][]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{
        buckets: make([][][]int, 1000),
    }
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    hash := key % 1000
    
    // target bucket is empty
    if len(this.buckets[hash]) == 0 {
        this.buckets[hash] = append(this.buckets[hash], []int{key, value})
        return 
    }
   
    // target bucket is not empty
    for index, item := range this.buckets[hash] {
        if item[0] == key {
            // key already in map
            this.buckets[hash][index][1] = value
            return
        }
    }
    
    // key not in map
    this.buckets[hash] = append(this.buckets[hash], []int{key, value})
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    hash := key % 1000
    for _, item := range this.buckets[hash] {
        if item[0] == key {
            return item[1]
        }
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    hash := key % 1000
    index := -1
    for i, item := range this.buckets[hash] {
        if item[0] == key {
            index = i
            break
        }
    }
    if index != -1 {
        this.buckets[hash] = append(this.buckets[hash][:index], this.buckets[hash][index+1:]...)
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
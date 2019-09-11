/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
type LRUCache struct {
	i int
	cap int
	cache map[int]int
	keys []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		i: 0,
		cap: capacity,
		cache: make(map[int]int),
		keys: nil
	}
}


func (this *LRUCache) Get(key int) int {
    if value, exists := this.cache[key]; exists {
		this.i++
		this.keys = append(this.keys, key)
		return value
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if _, exists := this.cache[key]; !exists {
		if len(this.cache) == this.cap {
			delete(this.cache, this.i - this.cap)
		}
	}
	this.i++
	this.cache[key] = value
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Hard (22.77%)
 * Total Accepted:    235.1K
 * Total Submissions: 1M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 * 
 * 
 * 
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 * 
 * 
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 * 
 * Example:
 * 
 * LRUCache cache = new LRUCache( 2 );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 * 
 * 
 */
 type LRUCache struct {
    cache map[int]int
	count map[int]int
	ts map[int]int64

    capacity int
    leastIndex int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: make(map[int]int, capacity),
        count: make(map[int]int, capacity),
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if v, ok := this.cache[key], ok {
		this.count[key]++
        return v
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if len(this.cache) == this.capacity {
		minUsedKey := 0
		minUsedKeyCount := 0
		minUsedCount := math.MaxInt32
		for k, v := range this.count {
			if v 
		}



        delete(this.count[minUsedKey])
        delete(this.cache[minUsedKey])
        this.cache[key] = value
        this.count[key] = 1
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

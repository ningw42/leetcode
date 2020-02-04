/*
 * @lc app=leetcode id=381 lang=golang
 *
 * [381] Insert Delete GetRandom O(1) - Duplicates allowed
 */

// @lc code=start
type RandomizedCollection struct {
	elems map[int][]int
	keys map[int]int
	count int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		elems: make(map[int][]int),
		keys: make(map[int]int),
		count: 0,
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	if _, exists := this.elems[val]; !exists {
		this.elems[val] = []int{this.count}
		this.keys[this.count] = val
		this.count++
		return true
	} else {
		this.elems[val] = append(this.elems[val], this.count)
		this.keys[this.count] = val
		this.count++
		return false
	}
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if keys, exists := this.elems[val]; exists && len(keys) > 0 {
		key := keys[0]
		this.elems[val] = this.elems[val][1:]
		delete(this.keys, key)
		return true
	} else {
		return false
	}
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	for {
		key := rand.Intn(this.count)
		if value, exists := this.keys[key]; exists {
			return value
		}
	}
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end


/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
    elems map[int]int
    keys map[int]int
    count int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        elems: make(map[int]int),
        keys: make(map[int]int),
        count: 0,
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.elems[val]; !exists {
        this.elems[val] = this.count
        this.keys[this.count] = val
        this.count++
        return true
    } else {
        return false
    }
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if key, exists := this.elems[val]; exists {
        delete(this.elems, val)
        delete(this.keys, key)
        return true
    } else {
        return false
    }
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    for {
        key := rand.Intn(this.count)
        if value, exists := this.keys[key]; exists {
            return value
        }
    }
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end


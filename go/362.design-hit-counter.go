type HitCounter struct {
    counter map[int]int
}


/** Initialize your data structure here. */
func Constructor() HitCounter {
    return HitCounter{
        counter: make(map[int]int),
    }
}


/** Record a hit.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int)  {
    this.counter[timestamp]++
}


/** Return the number of hits in the past 5 minutes.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
    count := 0
    for ts := timestamp; ts > timestamp - 300; ts-- {
        count += this.counter[ts]
    }
    return count
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
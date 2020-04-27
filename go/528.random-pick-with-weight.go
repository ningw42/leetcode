type Solution struct {
    w []int
    total int
}


func Constructor(w []int) Solution {
    total := 0
    for _, weight := range w {
        total += weight
    }
    return Solution{
        w: w,
        total: total,
    }
}


func (this *Solution) PickIndex() int {
    r := rand.Intn(this.total) + 1
    for i, weight := range this.w {
        r -= weight
        if r <= 0 {
            return i
        }
    }
    return len(this.w) - 1
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
type MovingAverage struct {
    cap int
    slice []int
    sum int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{cap: size}
}


func (this *MovingAverage) Next(val int) float64 {
    if len(this.slice) == this.cap {
        this.sum -= this.slice[0]
        this.slice = this.slice[1:]
        this.sum += val
        this.slice = append(this.slice, val)
        return float64(this.sum) / float64(this.cap)
    } else {
        this.sum += val
        this.slice = append(this.slice, val)
        return float64(this.sum) / float64(len(this.slice))
    }
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
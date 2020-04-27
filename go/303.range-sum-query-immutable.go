type NumArray struct {
	Nums []int
	Results map[int][]int
	Step int
}


func Constructor(nums []int) NumArray {
	results := make(map[int][]int)
	results[0] = nums
	return NumArray{
		Nums: nums,
		Results: results,
		Step: 0,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	step := j - i
	if step > this.Step {
		for s := this.Step + 1; s <= step; s++ {
			results := make([]int, len(this.Nums) - s)
			for begin := 0; begin + s < len(this.Nums); begin++ {
				results[begin] = this.Results[s-1][begin] + this.Nums[begin+s]
			}
			this.Results[s] = results
		}
		this.Step = step
	}
	return this.Results[step][i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
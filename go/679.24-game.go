// try out all possible ways
func judgePoint24(nums []int) bool {
    // cast to float64 to deal with division
    var candidates []float64
    for _, num := range nums {
        candidates = append(candidates, float64(num))
    }
    
    return tryout(candidates)
}

// if difference between result and 24 is smaller than epsilon
// result is considered equal to 24
const epsilon = 1e-4

// tryout, reduces two numbers into one
// by performing one of four operation on two choosen numbers
func tryout(nums []float64) bool {
    if len(nums) == 1 {
        return math.Abs(nums[0] - 24) < epsilon
    }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i == j {
                continue
            }
            // rest candidates
            var candidates []float64
            for k := 0; k < len(nums); k++ {
                if k == i || k == j {
                    continue
                }
                candidates = append(candidates, nums[k])
            }
            // all possible operations
            for _, reduced := range tryOperation(nums[i], nums[j]) {
                if tryout(append(candidates, reduced)) {
                    return true
                }
            }
        }
    }
    return false
}

// calculate fours possible operation bewteen a and b
func tryOperation(a, b float64) []float64 {
    return []float64{
        a + b,
        a - b,
        a * b,
        a / b,
    }
}
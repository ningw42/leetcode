func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }
    
    var segments []string
    var begin int
    var prev int
   
    begin = nums[0]
    prev = nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != prev + 1 {
            if begin == prev {
                segments = append(segments, strconv.Itoa(begin))
            } else {
                segments = append(segments, strconv.Itoa(begin) + "->" + strconv.Itoa(prev))
            }
            begin = nums[i]
        }
        prev = nums[i]
    }
    if begin == prev {
        segments = append(segments, strconv.Itoa(begin))
    } else {
        segments = append(segments, strconv.Itoa(begin) + "->" + strconv.Itoa(prev))
    }
   
    return segments
}
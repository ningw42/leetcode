func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    var i, j int
    intersection := make(map[int]struct{})
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            intersection[nums1[i]] = struct{}{}
            i++
            j++
        }
    }
    
    var slice []int
    for k, _ := range intersection {
        slice = append(slice, k)
    }
    
    return slice
}
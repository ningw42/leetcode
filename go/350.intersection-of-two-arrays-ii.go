func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    var i, j int
    var intersection []int
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            intersection = append(intersection, nums1[i])
            i++
            j++
        }
    }
    
    return intersection
}
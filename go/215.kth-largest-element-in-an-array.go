func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, 0, len(nums) - 1, k-1)
}

func quickSelect(nums []int, left, right, k int) int {
    if left == right {
        return nums[left]
    }
    pivotIndex := partition(nums, left, right)
    if pivotIndex > k {
        return quickSelect(nums, left, pivotIndex - 1, k)
    } else if pivotIndex < k {
        return quickSelect(nums, pivotIndex + 1, right, k)
    } else {
        return nums[k]
    }
}

func partition(nums []int, left, right int) int {
    pivotIndex := medianOfThree(nums, left, right)
    pivotValue := nums[pivotIndex]
    // swap nums[pivotIndex] and nums[right]
    swap(nums, pivotIndex, right)
    // select num that > pivotValue
    currentIndex := left
    for i := left; i < right; i++ {
        if nums[i] > pivotValue {
            // swap nums[i] and nums[currentIndex]
            swap(nums, currentIndex, i)
            currentIndex++
        }
    }
    // restore pivot's location
    swap(nums, currentIndex, right)
    return currentIndex
}

func medianOfThree(nums []int, left, right int) int {
    mid := (left + right) / 2
    // sort nums[left], nums[mid], nums[right]
    if nums[left] < nums[right] {
        swap(nums, left, right)
    }
    if nums[mid] < nums[right] {
        swap(nums, mid, right)
    }
    if nums[left] < nums[mid] {
        swap(nums, left, mid)
    }
    fmt.Println(nums[left], nums[mid], nums[right])
    return mid
}

func swap(nums []int, i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}
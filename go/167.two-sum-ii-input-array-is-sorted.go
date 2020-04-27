func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers) - 1
    for i < j {
        sum := numbers[i] + numbers[j]
        if sum < target {
            for i < j && numbers[i] == numbers[i+1] {i++}
            i++
        } else if sum > target {
            for i < j && numbers[j] == numbers[j-1] {j--}
            j--
        } else {
            return []int{i+1, j+1}
        }
    }
    return nil
}
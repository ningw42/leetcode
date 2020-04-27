func largestNumber(nums []int) string {
    var list []string
    var zero bool
    zero = true
    for _, num := range nums {
        list = append(list, strconv.Itoa(num))
        if num != 0 {
            zero = false
        }
    }
    if zero {
        return "0"
    } else {
        sort.Slice(list, getComparator(list))
        return strings.Join(list, "")
    }
}

// O(Nlog(N))
// sort by the first char
func getComparator(s []string) func(i, j int) bool {
    return func(i, j int) bool {
        ii := s[i][0]
        jj := s[j][0]
        if ii > jj {
            return true
        } else if ii < jj {
            return false
        } else {
            // first char of i and j are same
            return s[i] + s[j] > s[j] + s[i]
        }
    }
}

func nextGreaterElement(n int) int {
    result = nil
    first = true
    // find out all digits
    num, digits := countDigit(n)
    // match state when doing backtracking
    state := findState(num, 0, digits, map[int]struct{}{})
    // continue from the found state
    next(state, digits, 0, nil, map[int]struct{}{})
    // handle result
    if len(result) == 2 {
        element := 0
        for _, digit := range result[1] {
            element = element * 10 + digit
        }
        if element > math.MaxInt32 {
            return -1
        } else {
            return element
        }
    } else {
        return -1
    }
}

// result holds up to 2 permutation
// result[0] is identical to the given one
// result[1] is the next one
var result [][]int

// first indicates if next found permutation is exactly the given one
// controls the begin index of our backtracking
var first bool

func next(state []int, digits []int, index int, current []int, choosen map[int]struct{}) bool {
    if len(current) == len(digits) {
        c := make([]int, len(digits))
        copy(c, current)
        result = append(result, c)
        first = false
        return len(result) >= 2
    } else {
        var begin int
        if first {
            begin = state[index]
        } else {
            begin = 0
        }
        for i := begin; i < len(digits); i++ {
            if _, c := choosen[i]; !c {
                choosen[i] = struct{}{}
                if next(state, digits, index + 1, append(current, digits[i]), choosen) {
                    return true
                }
                delete(choosen, i)
                for i < len(digits) - 1 && digits[i] == digits[i+1] {i++}
            }
        }
        return false
    }
}

func findState(num []int, index int, digits []int, choosenDigitIndex map[int]struct{}) []int {
    if len(num) == index {
        return nil
    } else {
        for i := 0; i < len(digits); i++ {
            if _, choosen := choosenDigitIndex[i]; !choosen && num[index] == digits[i] {
                choosenDigitIndex[i] = struct{}{}
                return append([]int{i}, findState(num, index + 1, digits, choosenDigitIndex)...)
            }
        }
        // impossible to reach
        return nil
    }
}

func countDigit(n int) ([]int, []int) {
    var num []int
    for n != 0 {
        num = append(num, n % 10)
        n = n / 10
    }
    digits := make([]int, len(num))
    copy(digits, num)
    sort.Ints(digits)
    reverse(num)
    return num, digits
}

func reverse(s []int) {
    for i := 0; i < len(s) / 2; i++ {
        j := len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}
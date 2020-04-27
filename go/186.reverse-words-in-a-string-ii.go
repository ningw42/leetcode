func reverseWords(s []byte) {
    doubleReverse(s)
}

// 1. reverse entire s
// 2. reverse each word
func doubleReverse(s []byte) {
    reverse(s)
    begin := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            reverse(s[begin:i])
            begin = i + 1
        }
    }
    reverse(s[begin:])
}

func reverse(s []byte) {
    var i, j int
    for i = 0; i < len(s) / 2; i++ {
        j = len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}

// shift with ringbuffer, AC but quite slow
func ringbuffer(s []byte) {
    if len(s) == 0 {
        return 
    }
    var i, j int
    for i = 0; i < len(s) && s[i] != ' '; i++ {}
    for j = len(s) - 1; j >= 0 && s[j] != ' '; j-- {}
    if i > j {
        // single word in s
        // nothing needs to be done
        return 
    } else if i == j {
        // two words in s
    } else {
        // multiple words in s
        // reverse sub-problem first
        reverseWords(s[i+1:j])
    }
    // swap position by shift in ringbuffer
    // 1. shift left with length of left word in s
    shift(s, Left, i)
    // 2. shift right with length of right word in subarray of s that ends with right word
    shift(s[:len(s)-i], Right, len(s) - 1 - j)
}
const Left = true
const Right = false
// regard s as ringbuffer
// shift ringbuffer with direction and step count
// shift step by step since we dont have any extra space
func shift(s []byte, direction bool, step int) {
    for step > 0 {
        if direction == Left {
            stash := s[0]
            for k := 0; k < len(s) - 1; k++ {s[k] = s[k+1]}
            s[len(s)-1] = stash
        } else {
            stash := s[len(s)-1]
            for k := len(s) - 1; k > 0; k-- {s[k] = s[k-1]}
            s[0] = stash
        }
        step--
    }
}

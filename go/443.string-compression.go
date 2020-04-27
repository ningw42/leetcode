func compress(chars []byte) int {
    read := 0
    write := 0
    count := 0
    for read < len(chars) {
        count = 0
        for read < len(chars) - 1 && chars[read] == chars[read+1] {
            read++
            count++
        }
        chars[write] = chars[read]
        if count != 0 {
            for _, char := range []byte(strconv.Itoa(count+1)) {
                write++
                chars[write] = char
            }
        }
        read++
        write++
    }
    return write
}
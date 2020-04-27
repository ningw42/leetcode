func rand10() int {
    for {
        row, col := rand7(), rand7()
        if index := (row - 1) * 7 + col; index > 40 {
            continue
        } else {
            if result := index % 10; result == 0 {
                return 10
            } else {
                return result
            }
        }
    }
    return 0
}
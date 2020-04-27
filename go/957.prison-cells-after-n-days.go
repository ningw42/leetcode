// cells.length == 8
// cells[i] is in {0, 1}
// 1 <= N <= 10^9

func prisonAfterNDays(cells []int, N int) []int {
    current := toBitmap(cells)
    current = nextState(current)
    if N == 1 {
        return toSlice(current)
    }
   
    var recurrentStates []uint8
    recurrentStateIDs := make([]bool, 256)
    
    i := 1
    for i < N {
        if recurrent := recurrentStateIDs[int(current)]; recurrent {
            break
        }
        recurrentStateIDs[int(current)] = true
        recurrentStates = append(recurrentStates, current)
        current = nextState(current)
        i++
    }
    
    if i == N {
        return toSlice(current)
    } else {
        return toSlice(recurrentStates[(N-1)%len(recurrentStates)])
    }
}

func toBitmap(cells []int) uint8 {
    var state uint8
    for i := 0; i < len(cells); i++ {
        if cells[i] == 1 {
            state = state | (1 << (7-i))
        }
    }
    return state
}

func toSlice(state uint8) []int {
    var mask uint8 
    var s []int
    for mask = 0b1000_0000; mask > 0; mask = mask >> 1 {
        if (mask & state) == 0 {
            s = append(s, 0)
        } else {
            s = append(s, 1)
        }
    }
    return s
}

func nextState(current uint8) uint8 {
    var next, mask uint8
    for mask = 0b0000_0010; mask <= 0b0100_0000; mask = mask << 1 {
        left := ((mask << 1) & current) == 0
        right := ((mask >> 1) & current) == 0
        if left == right {
            next = next | mask
        }
    }
    return next
}
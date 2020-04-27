func fractionToDecimal(numerator int, denominator int) string {
    // divided by zero
    if denominator == 0 {
        return "0"
    }
   
    // prevent from int32 overflow
    // n for numerator
    // d for denominator
    var n, d uint64
    
    // sign
    var negative bool
    if numerator < 0 && denominator > 0 {
        n = uint64(-numerator)
        d = uint64(denominator)
        negative = true
    } else if numerator > 0 && denominator < 0 {
        n = uint64(numerator)
        d = uint64(-denominator)
        negative = true
    } else if numerator < 0 && denominator < 0 {
        n = uint64(-numerator)
        d = uint64(-denominator)
    } else {
        n = uint64(numerator)
        d = uint64(denominator)
    }
   
    // init
    integer := n / d
    n = n % d
    // integer
    if n == 0 {
        if negative {
            return fmt.Sprintf("-%d", integer)
        } else {
            return fmt.Sprintf("%d", integer)
        }
    }
    
    historyNumerator := make(map[uint64]uint64)
    var fraction []uint64
    var q, r, i uint64
    var result string
    for {
        n = n * 10
        // return if recurrent numerator occurs
        if begin, recurrent := historyNumerator[n]; recurrent {
            var recurrentPart, nonRecurrentPart strings.Builder
            for j := 0; j < int(begin); j++ {
                nonRecurrentPart.WriteString(fmt.Sprint(fraction[j]))
            }
            for j := int(begin); j < len(fraction); j++ {
                recurrentPart.WriteString(fmt.Sprint(fraction[j]))
            }
            result = fmt.Sprintf("%d.%s(%s)", integer, nonRecurrentPart.String(), recurrentPart.String())
            break
        } 
        // record current numerator
        historyNumerator[n] = i
        // does the math
        q = n / d
        r = n % d
        fraction = append(fraction, q)
        n = r
        if r == 0 {
            if len(fraction) == 0 {
                result = fmt.Sprint(integer)
            } else {
                var builder strings.Builder
                for _, f := range fraction {
                    builder.WriteString(fmt.Sprint(f))
                }
                result = fmt.Sprintf("%d.%s", integer, builder.String())
            }
            break
        }
        i++
    }
    
    // append sign
    if negative {
        return "-" + result
    } else {
        return result
    }
}
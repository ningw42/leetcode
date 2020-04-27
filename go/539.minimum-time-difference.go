func findMinDifference(timePoints []string) int {
    sort.Strings(timePoints)
    min := diff(timePoints[0], timePoints[len(timePoints)-1])
    for i := 0; i < len(timePoints)-1; i++ {
        if d := diff(timePoints[i], timePoints[i+1]); d < min {
            min = d
        }
    }
    return min
}

func diff(a, b string) int {
    if a < b {
        a, b = b, a
    }
    hour := (int(a[0]) - int(b[0])) * 10 + int(a[1]) - int(b[1])
    minute := (int(a[3]) - int(b[3])) * 10 + int(a[4]) - int(b[4])
    d := hour * 60 + minute
    return min(d, 1440-d)
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
func dayOfTheWeek(day int, month int, year int) string {
    from, to, before := sort([]int{year, month, day}, []int{2020, 4, 9})
    
    // day count for years
    dayCount := 0
    for y := from[0]; y < to[0]; y++ {
        dayCount += 365
        if isLeapYear(y) {
            dayCount++
        }
    }
    
    dayCount = dayCount - dayFromYearBegin(from[0], from[1], from[2]) + dayFromYearBegin(to[0], to[1], to[2])
    var weekday int
    if before {
        weekday = ((-dayCount) % 7 + 7) % 7
    } else {
        weekday = dayCount % 7
    }
    
    weekdays := []string{
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
    }
    
    return weekdays[weekday]
}

func sort(date1, date2 []int) ([]int, []int, bool) {
    if date1[0] < date2[0] {
        return date1, date2, true
    } else if date1[0] > date2[0] {
        return date2, date1, false
    } else {
        if date1[1] < date2[1] {
            return date1, date2, true
        } else if date1[1] > date2[1] {
            return date2, date1, false
        } else {
            if date1[2] < date2[2] {
                return date1, date2, true
            } else {
                return date2, date1, false
            }
        }
    }
}

func dayFromYearBegin(year, month, day int) int {
    counts := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    if isLeapYear(year) {
        counts[2] = 29
    }
    count := 0
    for i := 1; i < month; i++ {
        count += counts[i]
    }
    return count + day
}

func isLeapYear(year int) bool {
    if year % 4 != 0 {
        return false
    } else {
        if year % 100 != 0 {
            return true
        } else {
            if year % 400 == 0 {
                return true
            } else {
                return false
            }
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
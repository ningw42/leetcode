/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */

func employeeFreeTime(schedule [][]*Interval) []*Interval {
    return naive(schedule)
}

func naive(schedule [][]*Interval) []*Interval {
    var intervals []*Interval
    for _, p := range schedule {
        for _, interval := range p {
            intervals = append(intervals, interval)
        }
    }
    sort.Slice(intervals, func (i, j int) bool {
        ii, jj := intervals[i], intervals[j]
        if ii.Start < jj.Start {
            return true
        } else if ii.Start > jj.Start {
            return false
        } else {
            return ii.End < jj.End
        }
    })
    
    // printIntervals(intervals)
    var merged, free []*Interval
   
    var i, j, start, end int
    start = intervals[0].Start
    end = intervals[0].End
    for i < len(intervals) {
        for j < len(intervals) && intervals[j].Start == intervals[i].Start {
            j++
        }
        if j < len(intervals) {
            end = max(end, intervals[j-1].End)
            if intervals[j].Start > end {
                merged = append(merged, &Interval{
                    Start: start,
                    End: end,
                })
                start = intervals[j].Start
                end = intervals[j].End
            }
        }
        i = j
    }
    merged = append(merged, &Interval{
        Start: start,
        End: end,
    })
    
    // printIntervals(merged)
    
    for i := 0; i < len(merged) - 1; i++ {
        free = append(free, &Interval{
            Start: merged[i].End,
            End: merged[i+1].Start,
        })
    }
    
    return free
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func printIntervals(intervals []*Interval) {
    for i, interval := range intervals {
        fmt.Printf("[%d,%d]", interval.Start, interval.End)
        if i != len(intervals) - 1 {
            fmt.Print(",")
        }
    }
    fmt.Println()
}
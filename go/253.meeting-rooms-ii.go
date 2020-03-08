func minMeetingRooms(intervals [][]int) int {
    // sort meetings by start time asc
    sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
    // min heap to store rooms currently in use
    roomInUse := NewPQ()

    var roomsCount int
    for _, interval := range intervals {
        if roomsCount == 0 {
			// the very first iteration
            heap.Push(roomInUse, interval)
            roomsCount++
        } else {
            if roomInUse.Len() < roomsCount {
				// there are rooms available
                heap.Push(roomInUse, interval)
            } else {
				// there is no room available
                firstRoomToRelease := (*roomInUse)[0]
                if firstRoomToRelease[1] > interval[0] {
					// all rooms are in use
                    heap.Push(roomInUse, interval)
                    roomsCount++
                } else {
					// some rooms are not in use
                    for firstRoomToRelease[1] <= interval[0] {
                        heap.Pop(roomInUse)
                        if roomInUse.Len() == 0 {
                            break
                        }
                        firstRoomToRelease = (*roomInUse)[0]
                    }
                    heap.Push(roomInUse, interval)
                }
            }
        }
    }
    
    return roomsCount
}

type PQ [][]int
func NewPQ() *PQ {
    var container [][]int
    pq := PQ(container)
    return &pq
}
func (pq *PQ) Len() int {return len(*pq)}
func (pq *PQ) Swap(i, j int) {(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]}
func (pq *PQ) Less(i, j int) bool {return (*pq)[i][1] < (*pq)[j][1]}
func (pq *PQ) Push(item interface{}) {
    (*pq) = append((*pq), item.([]int))
}
func (pq *PQ) Pop() interface{} {
    length := pq.Len()
    item := (*pq)[length-1]
    (*pq) = (*pq)[:length-1]
    return item
}
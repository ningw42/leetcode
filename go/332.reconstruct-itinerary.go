/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */
func findItinerary(tickets [][]string) []string {
	// construct adjacency matrix
	matrix := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if _, exists := matrix[from]; !exists {
			matrix[from] = make([]string, 0)
		}
		matrix[from] = append(matrix[from], to)
	}

	// sort destinations by lexical order for every airport
	for _, destinations := range matrix {
		sort.Strings(destinations)
	}

	// DFS with backtracking
	// we are emulating the flight path, but add to 'itinerary' in the reverse order
	var itinerary, stack []string
	stack = append(stack, "JFK")
	for len(stack) > 0 {
		// top
		from := stack[len(stack)-1]
		if len(matrix[from]) == 0 {
			// no outgoing edge
			itinerary = append(itinerary, from)
			// pop
			stack = stack[:len(stack)-1]
		} else {
			// push
			stack = append(stack, matrix[from][0])
			matrix[from] = matrix[from][1:]
		}
	}

	// reverse itinerary
	for i, j := 0, len(itinerary)-1; i < j; i, j = i+1, j-1 {
		itinerary[i], itinerary[j] = itinerary[j], itinerary[i]
	}

	return itinerary
}

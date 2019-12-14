/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
// copied from 126.word ladder ii
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// misc preparation
	var numberOfVertex, begin, end int
	var beginInList, endInList bool
	for i := 0; i < len(wordList); i++ {
		if beginWord == wordList[i] {
			begin = i
			beginInList = true 
		}
		if endWord == wordList[i] {
			end = i
			endInList = true
		}
	}
	if !endInList {
		return 0
	}
	if beginInList {
		numberOfVertex = len(wordList)
	} else {
		begin = len(wordList)
		numberOfVertex = len(wordList)+1
	}
	// build graph
	var g [][]bool 
	for i := 0; i < numberOfVertex; i++ {
		g = append(g, make([]bool, numberOfVertex))
	}
	for i := 0; i < len(wordList); i++ {
		for j := i+1; j < len(wordList); j++ {
			if diff(wordList[i], wordList[j]) == 1 {
				g[i][j] = true
				g[j][i] = true
			}
		}
	}
	if !beginInList {
		for i := 0; i < len(wordList); i++ {
			if diff(beginWord, wordList[i]) == 1 {
				g[begin][i] = true
				g[i][begin] = true
			}
		}
	}

	// shortest path from begin to end
	// BFS
	queue := []int{begin}
	// seen holds if the vertex has been seen
	// seen[i] holds the shortest pathes from begin to i(share the same length)
	seen := make([][][]int, numberOfVertex)
	seen[begin] = [][]int{[]int{begin}}
	BFS:
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		for to, edgeExists := range g[from] {
			if edgeExists {
				// remove edge from graph
				g[from][to] = false
				g[to][from] = false

				if len(seen[to]) == 0 {
					// if vertex has not been seen yet
					// add to queue and mark as seen
					queue = append(queue, to)
					seen[to] = append(seen[to], appendToAll(seen[from], to)...)
				} else {
					// if vertex has been seen
					// if the new path share the same length, append 
					// if the new path length is greater than existing path length, ignore, if the 'to' vertex is our target, stop BFS
					// it is impossible that the length of new path is smaller than existing path length, since we are doing BFS
					if existingPathLength := len(seen[to][0]); existingPathLength == len(seen[from][0])+1 {
						seen[to] = append(seen[to], appendToAll(seen[from], to)...)
					} else if existingPathLength < len(seen[from][0])+1 {
						if to == end {
							break BFS
						}
					}
				}
			}
		}
	}

	// translate result
	if seen[end] == nil {
		return 0
	} else {
		return len(seen[end][0])
	}
}

func diff(a, b string) int {
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func appendToAll(s [][]int, t int) [][]int {
	var ret [][]int
	for _, p := range s {
		ret = append(ret, append(makeSliceCopy(p), t))
	}
	return ret
}

func makeSliceCopy(s []int) []int {
	copy := make([]int, len(s))
	for i, e := range s {
		copy[i] = e
	}
	return copy
}
// @lc code=end


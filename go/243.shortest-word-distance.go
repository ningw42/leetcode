func shortestDistance(words []string, word1 string, word2 string) int {
    index := make(map[string][]int)
    for i, word := range words {
        index[word] = append(index[word], i)
    }
    
    index1 := index[word1]
    index2 := index[word2]
   
    minDist := math.MaxInt32
    var i, j int
    for i < len(index1) && j < len(index2) {
        if index1[i] < index2[j] {
            minDist = min(minDist, index2[j]-index1[i])
            i++
        } else {
            minDist = min(minDist, index1[i]-index2[j])
            j++
        }
    }
    
    return minDist
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
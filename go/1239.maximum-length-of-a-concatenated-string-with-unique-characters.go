func maxLength(arr []string) int {
    return bruteForce(arr)
}

func bruteForce(arr []string) int {
    // build chars set for each string
    var hss []HashSet
    for _, s := range arr {
        hs := newHashSet(s)
        hss = append(hss, hs)
    }
  
    // calculate if any two string pair share common char
    intersected := make([][]bool, len(arr))
    for i := 0; i < len(arr); i++ {
        intersected[i] = make([]bool, len(arr))
    }
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if hss[i].intersect(hss[j]) {
                intersected[i][j] = true
                intersected[j][i] = true
            }
        }
    }

    // find all subsets of non-selfintersected elements of arr, and find the subset with max length
    var elements []int
    for i, hs := range hss {
        if !hs.selfIntersect() {
            elements = append(elements, i)
        }
    }
    max := 0
    subsets := findAllSubSets(elements)
    for _, subset := range subsets {
        subsetNotIntersects := true
        loop:
        for i := 0; i < len(subset); i++ {
            for j := i + 1; j < len(subset); j++ {
                if intersected[subset[i]][subset[j]] {
                    subsetNotIntersects = false
                    break loop
                }
            }
        }
        if subsetNotIntersects {
            length := 0
            for i := 0; i < len(subset); i++ {
                length += len(arr[subset[i]])
            }
            if length > max {
                max = length
            }
        }
    }
    
    return max
}

var allSubSets [][]int

func findAllSubSets(elements []int) [][]int {
    allSubSets = [][]int{}
    findSubSets(elements, 0, []int{})
    return allSubSets
}

func findSubSets(elements []int, i int, subset []int) {
    if i == len(elements) {
        allSubSets = append(allSubSets, makeSliceCopy(subset))
    } else {
        for j := i; j < len(elements); j++ {
            findSubSets(elements, j+1, append(subset, elements[j]))
        }
        findSubSets(elements, len(elements), subset)
    }
}

func makeSliceCopy(src []int) []int {
    dst := make([]int, len(src))
    copy(dst, src)
    return dst
}


type HashSet map[byte]int

func newHashSet(s string) HashSet {
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    return HashSet(m)
}

func (h HashSet) selfIntersect() bool {
    for _, count := range h {
        if count > 1 {
            return true
        }
    }
    return false
}

func (a HashSet) intersect(b HashSet) bool {
    for byteA := range a {
        if _, intersects := b[byteA]; intersects {
            return true
        }
    }
    return false
}
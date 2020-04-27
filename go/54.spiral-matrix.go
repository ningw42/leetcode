
const GoRight int = 0
const GoDown int = 1
const GoLeft int = 2
const GoUp int = 3

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}

	return traversal(matrix, 0, 0, GoRight)
}

func traversal(matrix [][]int, i, j, direction int) []int {
	// fmt.Println(i, j, direction)
	dirs := [][]int{
		[]int{i+1,j},
		[]int{i-1,j},
		[]int{i,j+1},
		[]int{i,j-1},
	}
	end := true
	for _, dir := range dirs {
		x := dir[0]
		y := dir[1]
		if 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0]) && matrix[x][y] != math.MaxInt32 {
			end = false
			break
		}
	}
	if end {
		last := matrix[i][j]
		matrix[i][j] = math.MaxInt32
		return []int{last}
	}

	var coods [][]int
	x := i
	y := j
	switch direction {
	case GoRight:
		for ; y < len(matrix[0]) && matrix[x][y] != math.MaxInt32; y++ {
			coods = append(coods, []int{x, y})
		}
		y--
	case GoDown:
		for ; x < len(matrix) && matrix[x][y] != math.MaxInt32; x++ {
			coods = append(coods, []int{x, y})
		}
		x--
	case GoLeft:
		for ; y >= 0 && matrix[x][y] != math.MaxInt32; y-- {
			coods = append(coods, []int{x, y})
		}
		y++
	case GoUp:
		for ; x >= 0 && matrix[x][y] != math.MaxInt32; x-- {
			coods = append(coods, []int{x, y})
		}
		x++
	}

	coods = coods[:len(coods)-1]
	// fmt.Println(coods)

	var row []int
	for _, cood := range coods {
		row = append(row, matrix[cood[0]][cood[1]])
		matrix[cood[0]][cood[1]] = math.MaxInt32
	}

	return append(row, traversal(matrix, x, y, (direction+1)%4)...)
}
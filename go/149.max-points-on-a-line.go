func maxPoints(points [][]int) int {
	m := make(map[string]int)
	for _, point := range points {
		m[strconv.Itoa(point[0])+","+strconv.Itoa(point[1])]++
	}

	var filteredPoints [][]int
	var count []int
	for point, c := range m {
		count = append(count, c)
		cood := strings.Split(point, ",")
		x, _ := strconv.Atoi(cood[0])
		y, _ := strconv.Atoi(cood[1])
		filteredPoints = append(filteredPoints, []int{x, y})
	}

	return maxPointsHelper(filteredPoints, count)
}

func maxPointsHelper(points [][]int, count []int) int {
	if len(points) <= 2 {
		var sum int
		for _, c := range count {
			sum += c
		}
		return sum
	}
	var checked [][]bool
	for i := 0; i < len(points); i++ {
		checked = append(checked, make([]bool, len(points)))
	}
	max := 0
	for i := 0; i < len(points); i++ {
		for j := i+1; j < len(points); j++ {
			if !checked[i][j] {
				numberOfPoints := count[i] + count[j]
				for k := 0; k < len(points); k++ {
					if k == i || k == j {
						continue
					}
					var onSameLine bool
					if points[i][0] == points[j][0] {
						onSameLine = points[k][0] == points[i][0]
					} else if points[i][1] == points[j][1] {
						onSameLine = points[k][1] == points[i][1]
					} else {
						K := 100 * float64(points[j][1] - points[i][1]) / float64(points[j][0] - points[i][0])
						onSameLine = 100 * float64(points[k][1] - points[i][1]) / float64(points[k][0] - points[i][0]) == K
					}
					if onSameLine {
						numberOfPoints += count[k]
						checked[i][k] = true
						checked[k][i] = true
						checked[j][k] = true
						checked[k][j] = true
					}
				}
				if numberOfPoints > max {
					max = numberOfPoints
				}
			}
		}
	}
	return max
}
/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 */
func robotSim(commands []int, obstacles [][]int) int {
	X, Y := getIndex(obstacles)
	// N, W, S, E
	// directions := []int{0, 1, 2, 3}
	x, y := 0, 0
	toward := 0
	max := 0
    for _, command := range commands {
		switch command {
		case -1:
			// turn right
			toward = (toward + 1) % 4
		case -2:
			// turn left
			toward = (toward + 3) % 4
		default:
			// go straight
			switch toward {
			case 0, 2:
				to := y + (1 - toward) * command
				if ys, exists := X[x]; exists {
					from := y
					isOb := false
					if toward == 0 {
						for _, oby := range ys {
							if oby > from && oby <= to {
								y = oby - 1
								isOb = true
								break
							}
						}
						if !isOb {
							y = to
						}
					} else {
						for i := len(ys) - 1; i >= 0; i-- {
							if ys[i] < from && ys[i] >= to {
								y = ys[i] + 1
								isOb = true
								break
							}
						}
						if !isOb {
							y = to
						}
					}
				} else {
					y = to
				}
			case 1, 3:
				to := x + (2 - toward) * command
				if xs, exists := Y[y]; exists {
					isOb := false
					from := x
					if toward == 1 {
						for _, obx := range xs {
							if obx > from && obx <= to {
								isOb = true
								x = obx - 1
								break
							}
						}
						if !isOb {
							x = to
						}
					} else {
						for i := len(xs) - 1; i >= 0; i-- {
							if xs[i] < from && xs[i] >= to {
								isOb = true
								x = xs[i] + 1
								break
							}
						}
						if !isOb {
							x = to
						}
					}
				} else {
					x = to
				}
			}
		}

		if area := x * x + y * y; area > max {
			max = area
		}
	}

	return max
}

func getIndex(obstacles [][]int) (map[int][]int, map[int][]int) {
	X := make(map[int][]int)
	Y := make(map[int][]int)
	for _, p := range obstacles {
		x, y := p[0], p[1]
		X[x] = append(X[x], y)
		Y[y] = append(Y[y], x)
	}

	for _, ys := range X {
		sort.Ints(ys)
	}
	for _, xs := range Y {
		sort.Ints(xs)
	}

	return X, Y
}


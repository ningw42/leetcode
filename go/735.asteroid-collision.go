func asteroidCollision(asteroids []int) []int {
    var next []int
    var collision, last bool
    collision = true
    for len(asteroids) > 0 && collision {
        collision = false
        last = false
        for i := 0; i < len(asteroids) - 1; {
            if asteroids[i] > 0 && asteroids[i+1] < 0 {
                collision = true
                if i == len(asteroids) - 2 {
                    last = true
                }
                if sum := asteroids[i] + asteroids[i+1]; sum > 0 {
                    next = append(next, asteroids[i])
                } else if sum < 0 {
                    next = append(next, asteroids[i+1])
                }
                i += 2
            } else {
                next = append(next, asteroids[i])
                i++
            }
        }
        if !last {
            next = append(next, asteroids[len(asteroids)-1])
        }
        asteroids, next = next, nil
    }
    return asteroids
}
func maxSatisfied(customers []int, grumpy []int, X int) int {
    base := 0
    for i := 0; i < len(customers); i++ {
        if grumpy[i] == 0 {
            base += customers[i]
        }
    }

    plus := 0
    for i := 0; i < X; i++ {
        if grumpy[i] == 1 {
            plus += customers[i]
        }
    }
    maxPlus := plus
    for i := 0; i < len(grumpy) - X; i++ {
        if grumpy[i+X] == 1 {
            plus += customers[i+X]
        }
        if grumpy[i] == 1 {
            plus -= customers[i]
        }
        if plus > maxPlus {
            maxPlus = plus
        }
    }
    
    return base + maxPlus
}
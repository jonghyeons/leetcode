func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    n := len(customers)
    totalSatisfied := 0

    for i := 0; i < n; i++ {
        if grumpy[i] == 0 {
            totalSatisfied += customers[i]
        }
    }

    maxIncrease := 0
    currentIncrease := 0
    for i := 0; i < minutes; i++ {
        if grumpy[i] == 1 {
            currentIncrease += customers[i]
        }
    }
    maxIncrease = currentIncrease

    for i := minutes; i < n; i++ {
        if grumpy[i] == 1 {
            currentIncrease += customers[i]
        }
        if grumpy[i - minutes] == 1 {
            currentIncrease -= customers[i - minutes]
        }
        if currentIncrease > maxIncrease {
            maxIncrease = currentIncrease
        }
    }

    return totalSatisfied + maxIncrease
}
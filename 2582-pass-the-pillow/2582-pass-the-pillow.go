func passThePillow(n int, time int) int {
    roundTripTime := (n - 1) * 2
    effectiveTime := time % roundTripTime

    if effectiveTime < n {
        return effectiveTime + 1
    } else {
        return 2*n - effectiveTime - 1
    }
}
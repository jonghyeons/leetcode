func numRescueBoats(people []int, limit int) int {
    var res int
    sort.Ints(people)
    left, right := 0, len(people)-1
    
    for left <= right {
        if people[left] + people[right] <= limit {
            left++
        }
        right--
        res++
    }

    return res
}
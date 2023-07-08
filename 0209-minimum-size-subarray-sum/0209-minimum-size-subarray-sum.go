func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    minLength := n + 1
    sum := 0
    left := 0

    for right := 0; right < n; right++ {
        sum += nums[right]

        for sum >= target {
            if right-left+1 < minLength {
                minLength = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }

    if minLength == n+1 {
        return 0
    }

    return minLength
}
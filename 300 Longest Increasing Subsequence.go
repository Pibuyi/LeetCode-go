func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp, res := make([]int, len(nums)), 0
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }
    for i := 0; i < len(dp); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        res = max(res, dp[i])
    }
    return res
}
func max(i, j int) int {
    if i > j {
        return i 
    }
    return j
}

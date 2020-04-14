func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := 1; i < amount + 1; i++ {
        dp[i] = 0x3f3f3f3f
    }
    dp[0] = 0
    for _, coin := range coins {
        for i := coin; i < amount + 1; i++ {
            dp[i] = min(dp[i], dp[i - coin] + 1)
        }
    }
    if dp[amount] == 0x3f3f3f3f {
        return -1
    }
    return dp[amount]
}

func min(i, j int) int {
    if i < j {
        return i 
    }
    return j
}

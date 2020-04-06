func max(i, j int) int { 
    if i > j {
        return i 
    }
    return j
}
func stoneGameIII(stoneValue []int) string {
    dp := make([]int, 50010)
    sum := 0
    n := len(stoneValue)
    for i := n - 1; i >= 0; i-- {
        dp[i] = -0x3f3f3f3f
        sum += stoneValue[i]
        for j := 1; j <= 3; j++ {
            dp[i] = max(dp[i], sum - dp[i + j])
        }
    }
    if sum - dp[0] == dp[0] {
        return "Tie"
    } else if sum - dp[0] > dp[0] {
        return "Bob"
    }
    return "Alice"
}

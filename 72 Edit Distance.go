func min(i, j, k int) int {
    t := i
    if t > j {
        t = j
    }
    if t > k {
        t = k
    }
    return t
}
func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)
    if n2 == 0 {
        return n1
    }
    dp := make([][]int, n1 + 1)
    for i:= 0; i < n1 + 1; i++ {
        dp[i] = make([]int, n2 + 1)
    }
    for j := 1; j < n2 + 1; j++ {
        dp[0][j] = dp[0][j - 1] + 1
    }
    for i := 1; i < n1 + 1; i++ {
        dp[i][0] = dp[i - 1][0] + 1
    }
    for i := 1; i < n1 + 1; i++ {
        for j := 1; j < n2 + 1; j++ {
            if word1[i -1] == word2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1]
            } else {
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1], dp[i -1][j - 1]) + 1
            }
        }
    }
    return dp[n1][n2]
}

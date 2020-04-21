func numberOfSubarrays(nums []int, k int) int {
    //用来保存在上一个奇数前的偶数个数
    dp := make([]int, 0)
    cnt, res := 0, 0
    for _, v := range nums {
        cnt++
        if v % 2 == 1 {
            dp = append(dp, cnt)
            cnt = 0
        }
        if len(dp) >= k {
            res += dp[len(dp) - k]
        }
    }
    return res

}

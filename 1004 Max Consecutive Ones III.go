func longestOnes(A []int, K int) int {
    n := len(A)
    if n == 0 {
        return 0
    }
    if n <= K {
        return n
    }
    left, right := 0, 0
    cnt, res := 0, K
    for right < n {
        if A[right] == 0 {
            cnt++
        }
        for cnt > K {
            if A[left] == 0 {
                cnt --
            }
            left++
        }
        right++
        if res < right - left {
            res = right - left
        }
    }
    return res
}

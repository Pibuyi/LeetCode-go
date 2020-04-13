func longestOnes(A []int, K int) int {
    n := len(A)
    if n == 0 {
        return 0
    }
    if n <= K {
        return n
    }
    left, right := 0, 0
    num, res := 0, K
    for right < n {
        if A[right] == 0 {
            num++
        }
        for num > K {
            if A[left] == 0 {
                num --
            }
            left++
        }
        right++
        res = max(res, right - left)
    }
    return res
}

func max(i, j int)  int {
    if i > j {
        return i 
    }
    return j
}

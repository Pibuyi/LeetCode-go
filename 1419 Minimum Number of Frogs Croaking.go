func minNumberOfFrogs(croakOfFrogs string) int {
    check := 0
    for _, v := range croakOfFrogs {
        check ^= int(v)
    }
    if check % 116 != 0 {
        return -1
    }
    res := 0
    a := make([]int, 5)
    for _, v := range croakOfFrogs {
        switch v {
            case 'c':
                a[0]++
            case 'r':
                a[1]++
            case 'o':
                a[2]++
            case 'a':
                a[3]++
            case 'k':
                a[4]++
        }
        for i := 1; i < 5; i++ {
            if a[i] > a[i - 1] {
                return -1
            }
            res = max(res,a)
            if a[4] == 1 {
                a[0]--
                a[1]--
                a[2]--
                a[3]--
                a[4]--
            }
        }
    }
    return res
}

func max(res int,a []int)int {
    for _, v := range a {
        if v > res {
            res = v
        }
    }
    return res
    
}

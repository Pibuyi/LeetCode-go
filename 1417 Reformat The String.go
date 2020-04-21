func reformat(s string) string {
    n := len(s)
    snum, sstr := 0, 0
    num, str := make([]byte, 0), make([]byte, 0)
    for i := 0; i < n; i++ {
        if s[i] >= '0' && s[i] <= '9' {
            snum++
            num = append(num, s[i])
        } else {
            sstr++
            str = append(str, s[i])
        }
    }
    if snum - sstr > 1 || snum - sstr < -1 {
        return ""
    }
    if snum < sstr {
        num, str = str, num
    }
    res := make([]byte, 0)
    for i := 0; i < len(num); i++ {
        res = append(res, num[i])
        if i < len(str) {
            res = append(res, str[i])
        }
    }
    return string(res[:])
    
}

func restoreIpAddresses(s string) []string {
    var ip string
    res := make([]string, 0)
    helper(s, 0, ip, &res)
    return res
}

func helper(s string, n int, ip string, res *[]string) {
    if n == 4 {
        if s == "" {
            *res = append(*res, ip)
        }
    } else {
        for i := 1; i < 4; i++ {
            if len(s) < i {
                break
            }
            val, _ := strconv.Atoi(s[:i])
            //去掉值大于255和以0开头的不符合规定的地址
            if val > 255 || (s[0] == '0' && len(s[:i]) != 1) {
                continue
            }
            if n < 3 {
                helper(s[i:], n + 1, ip + s[:i] + ".", res)
            } else {
                helper(s[i:], n + 1, ip + s[:i], res)  
            }
        }
    }
}

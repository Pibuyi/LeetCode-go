func getHash(s string, len int) int {
    h := 1
    for i := 0; i < len; i++ {
        h *= int(s[i])
    }
    return h
}

func findAnagrams(s string, p string) []int {
    sLen := len(s)
    pLen := len(p)
    ans := []int{}
    pHash := getHash(p, pLen)
    for i := 0; i <= sLen - pLen; i++ {
        if getHash(s[i : i + pLen], pLen) == pHash {
            ans = append(ans, i)
        }
    }
    return ans
}

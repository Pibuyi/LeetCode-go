func CheckPermutation(s1 string, s2 string) bool {
    m :=  make([]int, 125)
    if len(s1) != len(s2) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        m[s1[i]]++
        m[s2[i]]--
    }
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    return true
}

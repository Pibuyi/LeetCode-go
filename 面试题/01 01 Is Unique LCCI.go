func isUnique(astr string) bool {
    m :=  make([]int, 125)
    for _, v := range astr {
        m[v]++
        if m[v] > 1 {
            return false
        }
    }
    return true
}

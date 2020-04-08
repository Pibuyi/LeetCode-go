func hIndex(citations []int) int {
    sort.Slice(citations, func(i, j int) bool{
        return citations[i] > citations[j]
    })
    res := 0
    for k, v := range citations {
        if v >= k + 1{
            res = k + 1
        } else {
            break
        }
    }
    return res
}

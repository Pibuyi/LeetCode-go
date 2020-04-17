func dailyTemperatures(T []int) []int {
    res := make([]int, 0)
    for i, v := range T {
        now := v
        j := i
        ok := false
        for j < len(T) {
            if T[j] > now {
                res = append(res, j - i)
                ok = true
                break
            }
            ok = false
            j++
        }
        if !ok {
            res = append(res, 0)
        }
    }
    return res
}

func processQueries(queries []int, m int) []int {
    res := make([]int, 0)
    p := make([]int, m)
    for i := 0; i < m; i++ {
        p[i] = i + 1
    }
    for _, q := range queries {
        res = append(res, update(q, p, m))
    }
    return res
}

func update(q int,p []int, m int) (t int) {
    for i := 0; i < m; i++ {
        if q == p[i] {
            t = i
            break
        }
    }
    temp := p[t]
    for i := t; i > 0; i-- {
        p[i] = p[i - 1]
    }
    p[0] = temp
    return t
}

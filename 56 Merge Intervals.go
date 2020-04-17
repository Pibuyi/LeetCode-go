func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    // fmt.Println(intervals)
    res := make([][]int, 0)
    a, b := 0, 0
    for i := 0; i < len(intervals); i++ {
        a, b = intervals[i][0], intervals[i][1]
        for i + 1 < len(intervals) && b >= intervals[i + 1][0] {
            i++
            if b < intervals[i][1] {
                b = intervals[i][1]
            }
        }
       
        res = append(res, []int{a, b})
    }
    return res
}

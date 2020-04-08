func reconstructQueue(people [][]int) [][]int {
    len := len(people)
    res := make([][]int, 0)
    // 排序 身高降序， 前面的升序
    sort.Slice(people, func(i, j int) bool{
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }
        return people[i][0] > people[j][0]
    })

    // 两种情况，如果前面的人数大于等于当前总人数，直接在后边插入，否则，在第k个位置插入
    for i := 0; i < len; i++ {
        pre := people[i][1]
        if pre >= len {
            res = append(res, people[i])
        } else {
            res = append(res, people[i])    // 申请空间
            // 插入位置
            copy(res[pre + 1:], res[pre:])
            res[pre] = people[i]
        }
    }
    return res
}


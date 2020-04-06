func max(i, j int) int {
    if i > j {
        return i 
    }
    return j
}
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    cur, pre := 0, 0
    for _, v := range nums {
        cur, pre = max(pre + v, cur), cur
    }
    pre = 0
    return cur
}

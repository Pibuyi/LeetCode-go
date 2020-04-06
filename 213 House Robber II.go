func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    return max(helper(nums[:len(nums) - 1]), helper(nums[1:]))
}

func helper(nums []int) int {
    pre, cur := 0, 0
    for _, v := range nums {
        cur, pre = max(pre + v, cur), cur
    }
    return cur
}

func max(i, j int) int {
    if i > j {
        return i 
    }
    return j
}

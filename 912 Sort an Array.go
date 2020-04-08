func sortArray(nums []int) []int {
    s := [100010] int{}
    for i := 0; i < len(nums); i++ {
        s[nums[i]+50000]++
    }
    index := 0 
    for i := 0; i < 100001; i++ {
        for s[i] > 0 {
            nums[index] = i - 50000
            index++
            s[i]--
        }
    }
    return nums
}

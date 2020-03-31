func calculate(nums []int, i int, sum int, S int,count *int){
    if i == len(nums) {
        if sum == S {
            *count = *count +  1
            // fmt.Println(*count)
        }
    } else {
        calculate(nums,i + 1, sum + nums[i], S, count)
        calculate(nums, i + 1, sum - nums[i], S, count)
    }
}
func findTargetSumWays(nums []int, S int) int {
    count := 0
    calculate(nums, 0, 0, S, &count)
    return count
}

func canJump(nums []int) bool {
    n := len(nums)
    t := 0
    for i := 0; i < n; i++ {
        if i > t {
            return false
        } else {
            if t < i + nums[i] {
                t = i + nums[i]
            }
        }
    }
    
    return true
}

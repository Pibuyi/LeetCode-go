func twoSum(nums []int, target int) []int {
	m := map[int]int{}
    for i, v := range nums {
        if k, ok := m[target - v]; ok {
            return []int{k, i}
        }
        m[v] = i
    }
    return nil
}
/*
func twoSum(nums []int, target int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
*/

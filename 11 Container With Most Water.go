func maxArea(height []int) int {
    left, right := 0, len(height) - 1
    res := 0
    for left < right {
        res = max(res, (right - left) * min(height[left], height[right]))
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return res
}

func max(i, j int) int {
    if i > j {
        return i 
    }
    return j
}
func min(i, j int) int {
    if i < j {
        return i 
    }
    return j
}

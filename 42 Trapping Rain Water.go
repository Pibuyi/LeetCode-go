func trap(height []int) int {
    res := 0
    Mleft, Mright := 0, 0
    left, right := 0, len(height) - 1
    for left < right {
        if height[left] < height[right] {
            if height[left] < Mleft {
                res += (Mleft - height[left])
            } else {
                Mleft = height[left]
            }
            left++
        } else {
            if height[right] < Mright {
                res += (Mright - height[right])
            } else {
                Mright = height[right]
            }
            right--
        }
    }
    return res
}

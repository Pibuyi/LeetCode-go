func smallestRepunitDivByK(K int) int {
    //爆范围
    if K % 2 == 0 || K % 5 == 0 {
        return -1
    }
    i, j := 1, 1
    for i % K != 0 {
        i %= K
        i = i *10 + 1
        j++
    }
    return j 
}

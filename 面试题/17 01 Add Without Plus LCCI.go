func add(a int, b int) int {
    // ^不同为1 &相同为1
    for b != 0 {
        a, b = (a ^ b), (a & b) << 1
    }
    return a
}

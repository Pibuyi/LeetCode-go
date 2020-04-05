func numSteps(s string) int {
    steps := 0
    // 1101 1110 111 1000 100 10 1
    for (len(s) != 1){     
        n := len(s)
        if s[n - 1] == '1' {
            steps++
            s = helper(s, n)   
        }
        steps++
        s = s[:len(s) - 1]
    }
    return steps
}
// 模拟字符串的变化，修改字符串
func helper(str  string, n int) string {
    s := []rune(str)
    i := n - 1
    for i = n - 1; i > 0; i-- {
        s[i] = '0'
        if s[i - 1] != '1' {
            s[i - 1] = '1'
            break
        }
        s[i - 1] = '1'
        
    }
    if i == 0 {
        s = append(s, '0') 
    }
    return string(s)
}

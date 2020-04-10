func reverseWords(s string) string {
    res := ""
    falg := false
    for i := len(s) - 1; i >= 0; i-- {
        for i >= 0 && s[i] == ' ' {
            i--
        }
        j := i
        for j >= 0 && s[j] != ' ' {
            j--
        }
        if falg && i != j {
            res += " "
        }
        res += s[j + 1: i + 1]
        i = j
        falg = true
    }
    return res
}

func toLowerCase(str string) string {
    var res []byte
    for i := 0; i < len(str); i++ {
        if str[i] >= 'A' && str[i] <= 'Z' {
            res = append(res, str[i] + 32)
        } else {
            res = append(res, str[i])
        }
    }
    return string(res)
}

//方法一
func myAtoi(str string) int {
    return convert(clean(str))
}
//将字符串转换成全数字
func clean(str string) (sign int, arr string) {
    str = strings.TrimSpace(str)
    if str == "" {
        return
    }
    switch str[0] {
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' :
            sign, arr = 1, str
        case '+' :
            sign, arr = 1, str[1:]
        case '-' :
            sign, arr = -1, str[1:]
        default :
            str = ""
            sign = 0
            return
    }
    for i, _ := range(arr) {
        if arr[i] < '0' || arr[i] > '9' {
            arr = arr[:i]
            break
        }
    }
    return
}

func convert(sign int, arr string) int {
    absNum := 0
    for _ , s := range(arr) {
        absNum = absNum*10 + int(s - '0')
        if sign == 1 && absNum * sign > math.MaxInt32 {
            return math.MaxInt32
        } else if sign == -1 && absNum * sign < math.MinInt32 {
            return math.MinInt32
        }
    }
    return sign * absNum
}

//方法二
func myAtoi2(str string) int {
    // 除前缀空格
    str = strings.TrimSpace(str)
    l := len(str)
    // 特判
    if l == 0 {
        return 0
    }
    if l == 1 && (str[0] < '0' || str[0] > '9') {
        return 0
    }
    
    
    i, ans := 0, 0
    // 进入循环条件
    if (str[i] == '-' || str[i] == '+') || (str[i] >= '0' && str[i] <= '9') {
        begin, end := i, i + 1
        for end < l && (str[end] >= '0' && str[end] <= '9') {
            end++
        }
        ans, _ = strconv.Atoi(str[begin: end])
    } else {
        return 0
    }

    if ans > math.MaxInt32 {
        return math.MaxInt32
    } else if ans < math.MinInt32 {
        return math.MinInt32
    }
    return ans
}

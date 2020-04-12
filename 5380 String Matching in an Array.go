func stringMatching(words []string) []string {
    sort.Slice(words, func(i, j int) bool{
        return len(words[i]) < len(words[j])
    })
    res := make([]string, 0)
    for i := len(words) - 1; i >= 1; i-- {
        for j := i - 1; j >= 0; j-- {
            if strings.Contains(words[i],words[j]) == true {
                res = append(res, words[j])
            }
        }
    }
    return RemoveRepeatedElement(res)
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
    newArr = make([]string, 0)
    for i := 0; i < len(arr); i++ {
        repeat := false
        for j := i + 1; j < len(arr); j++ {
            if arr[i] == arr[j] {
                repeat = true
                break
            }
        }
        if !repeat {
            newArr = append(newArr, arr[i])
        }
    }
    return
}

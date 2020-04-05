// type Interface interface {
//     // Len方法返回集合中的元素个数
//     Len() int
//     // Less方法报告索引i的元素是否比索引j的元素小
//     Less(i, j int) bool
//     // Swap方法交换索引i和j的两个元素
//     Swap(i, j int)
// }
// 排序算法需要实现的接口

type IntSlice []int
func (s IntSlice) Len() int {return len(s)}
func (s IntSlice) Swap(i, j int) {s[i], s[j] = s[j], s[i]}
func (s IntSlice) Less(i, j int)bool {return s[i] > s[j]}

func minSubsequence(nums []int) []int {
    // 计算数组总和
    sum := 0
    for _, v := range nums {
        sum += v
    }
    // 找出至少中间值
    temp := sum /2
    
    sort.Sort(IntSlice(nums))
    res := make([]int , 0)
    cur := 0
    for _, v := range nums {
        // 如果当前值小于或等于中间值，添加到答案中，否则退出循环
        if cur <= temp {
            cur += v
            res = append(res, v)
        } else {
            break
        }
    }
    return res
}

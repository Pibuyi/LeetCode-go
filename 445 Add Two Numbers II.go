/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1, s2 := make([]int, 0), make([]int, 0)
    for l1 != nil {
        s1 = append(s1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        s2 = append(s2, l2.Val)
        l2 = l2.Next
    }
    flag, sum := 0, 0
    var res *ListNode
    res = nil
    for len(s1) > 0 || len(s2) > 0 {
        a, b := 0, 0
        if len(s1) >= 1 {
            a = s1[len(s1) - 1]
            s1 = s1[:len(s1) - 1]
        }
        if len(s2) >= 1 {
            b = s2[len(s2) - 1]
            s2 = s2[:len(s2) - 1]
        }
        sum = (a + b + flag)
        flag = sum / 10
        sum %= 10
        var cur = &ListNode{Val: sum, Next: res}
        res = cur 
    }
    if flag > 0 {
        var cur = &ListNode{Val: flag, Next: res}
        res = cur 
    }
    return res
}

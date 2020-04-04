/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
   n := 0
   for p := head; p != nil; p = p.Next {
       n++
   }
   //计算出中点位置，从中点开始反转链表， n + 1为了处理节点个数是奇数的情况
   p := head
   for i := 0; i < (n + 1) / 2; i++ {
       p = p.Next
   }
   p = reverse(p)
   for p != nil {
       if p.Val != head.Val {
           return false
       }
       p, head = p.Next, head.Next
   }
   return true
}

func reverse(p *ListNode) *ListNode {
    var temp *ListNode
    for p != nil {
        p, p.Next, temp = p.Next, temp, p
    }
    return temp
}

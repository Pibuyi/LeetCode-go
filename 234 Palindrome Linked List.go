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
   p := head
   for i := 0; i < n / 2; i++ {
       p = p.Next
   }
   if n % 2 == 1 {
       p = p.Next
   }
   //从中点开始反转链表
   p = reverse(p)
   for p != nil {
       if head.Val != p.Val {
           return false
       }
       p, head = p.Next, head.Next
   }
   return true
}

func reverse(p *ListNode) *ListNode {
    var pre *ListNode
    for p != nil {
        p.Next, pre, p = pre, p, p.Next
    }
    return pre
}

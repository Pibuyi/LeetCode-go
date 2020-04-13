/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var  p *ListNode = nil
    for head != nil {
        p, head.Next, head = head, p,  head.Next  
    }
    return p
}

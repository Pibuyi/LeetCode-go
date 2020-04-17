/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    fast := head.Next
    for fast != nil && head != nil && fast.Next != nil {
        if fast == head {
            return true
        }
        fast = fast.Next.Next
        head = head.Next
    }
    return false
}

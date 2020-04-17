/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    cnt := 0
    if root.Val == sum {
        cnt++
    }
    cnt += helper(root.Left, sum - root.Val)
    cnt += helper(root.Right, sum - root.Val)
    return cnt
}

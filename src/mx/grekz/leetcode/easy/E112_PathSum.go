package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if ( root == nil ) {return false}
    sumOk := sum - root.Val == 0 && root.Left == nil && root.Right == nil
    remain := sum - root.Val
    return sumOk || hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
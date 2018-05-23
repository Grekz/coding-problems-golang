package medium

// @author grekz
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    doit(root, &res)
    return res
}
func doit(n *TreeNode, res *[]int){
    if n != nil {
        doit(n.Left,res)
        *res = append(*res, n.Val)
        doit(n.Right,res)
    }
}
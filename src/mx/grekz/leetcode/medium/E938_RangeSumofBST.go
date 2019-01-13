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
 func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil { return 0 }
    res := 0
    if L <= root.Val && root.Val <= R {
        res += root.Val
    }
    if L <= root.Val {
        res += rangeSumBST(root.Left, L, R)
    }
    if R >= root.Val {
        res += rangeSumBST(root.Right, L, R)
    }
    return res
}
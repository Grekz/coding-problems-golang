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
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) < 1 {
        return nil
    }
    i, m := maxIndex(nums)
    return &TreeNode{
        Val: m,
        Left: constructMaximumBinaryTree(nums[:i]),
        Right: constructMaximumBinaryTree(nums[i+1:]),
    }
}
func maxIndex(nums []int) (r, m int) {
    m = nums[0]
    for i, x := range nums {
        if x > m {
            r = i
            m = x
        }
    }
    return
}
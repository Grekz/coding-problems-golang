package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if ( len(nums) == 0 ) { return nil }
    return helper(nums, 0, len(nums)-1)
}
func helper( nums []int, l int, h int ) *TreeNode{
    if ( l > h ) { return nil }
    m := l + ( h - l ) / 2
    res := new(TreeNode)
    res.Val = nums[m]
    res.Left = helper(nums, l, m-1)
    res.Right = helper(nums, m+1, h)
    return res
}
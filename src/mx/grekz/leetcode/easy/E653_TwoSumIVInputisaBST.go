package easy

// @author grekz
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findTarget(root *TreeNode, k int) bool {
    s := make(map[int]bool)
    return dfs(root, k, s)    
}
func dfs(node *TreeNode, k int, s map[int]bool) bool {
    if node == nil {
        return false
    } 
    if _, ok := s[k - node.Val]; ok {
        return true
    }
    s[node.Val] = true
    return dfs(node.Left, k, s) || dfs(node.Right, k, s)
}
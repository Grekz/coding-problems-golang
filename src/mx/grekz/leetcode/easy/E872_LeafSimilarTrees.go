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
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return dfs(root1) == dfs(root2)
}
func dfs( n *TreeNode ) string {
    if n == nil {
        return ""
    }
    if n.Left == n.Right {
        return string(n.Val)
    }
    return dfs(n.Left) + dfs(n.Right)
}
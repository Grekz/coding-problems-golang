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
 var ans int
 func distributeCoins(root *TreeNode) int {
	 ans = 0
	 dfs(root)
	 return ans
 }
 func dfs(n *TreeNode) int {
	 if n == nil { return 0 }
	 l, r := dfs(n.Left), dfs(n.Right)
	 ans += abs(l) + abs(r)
	 return r + l + n.Val - 1
	 
 }
 func abs( a int ) int {
	 if a < 0 {
		 return -a
	 }
	 return a
 }
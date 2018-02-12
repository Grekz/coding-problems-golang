package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1;
}
func dfs(root *TreeNode) int{
    if ( root == nil ) { return 0; }
    lh := dfs(root.Left);
    if ( lh == -1 ) { return lh; }
    rh := dfs(root.Right);
    if ( rh == -1 ) { return rh; }
    dif := lh - rh;
    if ( dif > 1 || dif < -1 ) { return -1; }
    return Max(lh, rh) + 1 ;
}
func Max( a int, b int) int{
    if a > b {
        return a
    }
    return b
}
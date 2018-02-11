package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
    res := [][]int{}
    dfs(&res, 0, root)
    return res
}
func dfs(res *[][]int, lvl int, root *TreeNode) {
    if ( root == nil ) { 
        return 
    }
    le := len(*res)
    lvl += 1
    if ( le < lvl ) {
        *res = append([][]int{{}}, *res...)
        le+=1
    }
    (*res)[le - lvl] = append((*res)[le - lvl], root.Val)
    dfs(res, lvl, root.Left) 
    dfs(res, lvl, root.Right)
}
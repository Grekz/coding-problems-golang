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
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root != nil {
        doit(&res, root, 0)
    }
    return res
}
func doit(res *[][]int, walk *TreeNode, lvl int) {
    tmp := []int{}
    if len(*res) <= lvl {
        *res = append(*res, tmp)
    }else{
        tmp = (*res)[lvl]
    }
    tmp = append(tmp, walk.Val)
    (*res)[lvl] = tmp
    if walk.Left != nil {
        doit(res, walk.Left, lvl+1)
    }
    if walk.Right != nil {
        doit(res, walk.Right, lvl+1)
    }
}
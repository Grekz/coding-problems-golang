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
import "strconv"
func binaryTreePaths(root *TreeNode) []string {
    res := []string{}
    fmt.Print("started")
    if ( root != nil ) {
        traverse(root, "", &res)
    }
    return res
}

func traverse(node *TreeNode, cur string, res *[]string) {
    tmp := cur + strconv.Itoa(node.Val)
    if ( node.Right == nil && node.Left == nil ) {
        *res = append(*res, tmp)
    }else{
        tmp += "->"
        if ( node.Left != nil ) {
            traverse(node.Left, tmp, res)
        }
        if ( node.Right != nil ) {
            traverse(node.Right, tmp, res)
        }
    }
}
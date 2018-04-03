package easy

// @author grekz
func sumOfLeftLeaves(root *TreeNode) int {
    if (root == nil || ( root.Left == nil && root.Right == nil ) ) { return 0 } 
        
    ans := 0
    if(root.Left != nil && root.Left.Left == nil && root.Left.Right == nil) { ans += root.Left.Val } 
    return ans + sumOfLeftLeaves(root.Right) + sumOfLeftLeaves(root.Left)
}
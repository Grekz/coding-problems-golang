package easy

func maxSubArray(nums []int) int {
    var res, cur int = nums[0], nums[0]
    for i := 1; i < len(nums); i++{
        cur = Max(cur+nums[i], nums[i])
        res = Max(cur,res)
    }
    return res
}
func Max( a int, b int) int {
    if a > b {
        return a
    }
    return b
}
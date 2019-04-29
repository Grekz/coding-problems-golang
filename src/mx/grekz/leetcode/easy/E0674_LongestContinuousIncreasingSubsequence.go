package easy

// @author grekz
func findLengthOfLCIS(nums []int) int {
    res, count, n := 0, 0, len(nums)
    for i := 0; i < n; i++ {
        if i < 1 || nums[i-1] < nums[i] {
            count++
            res = max(res, count)
        } else {
            count = 1  
        }
    }
    return res
}
func max( a, b int ) int {
    if a > b {
        return a
    }
    return b
}
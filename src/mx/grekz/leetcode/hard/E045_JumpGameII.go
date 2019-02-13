package hard

// @author grekz
func jump(nums []int) int {
    count, end, far, n := 0, 0, 0, len(nums) - 1
    for i := 0; i < n; i++ {
        far = max(far, i + nums[i])
        if i == end {
            count++
            end = far
        }
    }
    return count
}
func max( a, b int) int {
    if a < b {
        return b
    }
    return a
}
package medium

// @author grekz
func findDuplicate(nums []int) int {
    n := len(nums)
    if n > 1 {
        for i := 0 ; i < n; i++ {
            cur := abs(nums[i]) - 1
            if nums[cur] < 0 {
                return cur + 1
            }
            nums[cur] = -nums[cur]
        }
    }
    return -1
}
func abs( a int ) int {
    if a < 0 {
        return -a
    }
    return a
}
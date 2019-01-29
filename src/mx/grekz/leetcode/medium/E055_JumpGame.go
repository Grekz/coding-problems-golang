package medium

// @author grekz
func canJump(nums []int) bool {
    last := len(nums) - 1
    for i := last; i >= 0; i-- {
        if i + nums[i] >= last {
            last = i
        }
    }
    return last == 0
}
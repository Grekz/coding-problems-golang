package easy

// @author grekz
func checkPossibility(nums []int) bool {
    n := len(nums)
    exists := false
    for i := 1; i < n; i++ {
        if nums[i] < nums[i-1] {
            if exists {
                return false
            }
            exists = true
            if i >= 2 && nums[i] < nums[i-2] {
                nums[i] = nums[i-1]
            }
        }
    }
    return true
}
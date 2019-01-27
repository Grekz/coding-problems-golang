package medium

// @author grekz
func nextPermutation(nums []int)  {
    n := len(nums) - 1
    i, j := n - 1, n
    for i >= 0 && nums[i] >= nums[i+1]{ i-- }
    if i >= 0 {
        for nums[j] <= nums[i] { j-- }
        nums[j], nums[i] = nums[i], nums[j]
    }
    i++
    for i < n {
        nums[n], nums[i] = nums[i], nums[n]
        n--
        i++
    }
}
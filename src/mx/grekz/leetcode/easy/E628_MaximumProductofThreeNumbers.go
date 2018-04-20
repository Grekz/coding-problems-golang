package easy

// @author grekz
import "sort"
func maximumProduct(nums []int) int {
    sort.Ints(nums)
    n := len(nums) - 1
    a, b := nums[n] * nums[n-1] * nums[n-2], nums[0] * nums[1] * nums[n]
    if a > b {
        return a  
    }
    return b
}
package medium

// @author grekz
func removeDuplicates(nums []int) int {
    if len(nums) < 3 { return len(nums) }
    i := 0 
    for _, x := range nums {
        if i < 2 || nums[i-2] < x {
            nums[i] = x
            i += 1        
        }
    }
    return i
}
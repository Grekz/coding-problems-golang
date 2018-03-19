package easy

// @author grekz
func missingNumber(nums []int) int {
    res := 0
    for i, e := range nums {
        res ^= i ^ e
    }
    return res ^ len(nums)
}
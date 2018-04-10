package easy

// @author grekz
func minMoves(nums []int) int {
    m := min(nums)
    res := 0
    for _, x := range nums {
        res += x - m
    }
    return res
}
func min(nums []int) int {
    r := 2147483647
    for _, x := range nums {
        if x < r {
            r = x
        }
    }
    return r
}
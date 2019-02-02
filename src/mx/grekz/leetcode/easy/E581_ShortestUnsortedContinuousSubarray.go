package easy

// @author grekz
func findUnsortedSubarray(nums []int) int {
    n, mi, ma, ini, end, b := len(nums) - 1, 2147483647, -2147483648, -1, -2, 0
    for i, a := range nums {
        b = nums[n - i]
        ma = max(a, ma)
        mi = min(b, mi)
        if a < ma {
            end = i
        }
        if b > mi {
            ini = n - i
        }
    }
    return end - ini + 1
}
func max( a, b int) int {
    if a > b {
        return a
    }
    return b
}
func min( a, b int) int {
    if a < b {
        return a
    }
    return b
}
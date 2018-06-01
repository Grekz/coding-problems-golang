package medium

// @author grekz
func productExceptSelf(nums []int) []int {
    p, n, t := make([]int, 0), len(nums) - 1, 1
    for i := range nums {
        p = append(p, t)
        t *= nums[i]
    }
    t = 1
    for k := range nums {
        i := n - k
        p[i] *= t
        t *= nums[i]
    }
    return p
}
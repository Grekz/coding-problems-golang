package easy

// @author grekz
func findDisappearedNumbers(nums []int) []int {
    res := make([]int, 0, len(nums) + 1)
    for _, x := range nums {
        v := abs(x) - 1
        nums[v] = neg(nums[v])
    }
    for i, x := range nums {
        if x > 0 {
            res = append(res, i + 1)
        }
    }
    return res
}
func abs(x int) int{
    if x < 0 {
        return -x
    }
    return x
}
func neg(x int) int{
    if x > 0 {
        return -x
    }
    return x
}
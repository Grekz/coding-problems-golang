package easy

// @author grekz
import "sort"
func arrayPairSum(nums []int) int {
    res := 0
    sort.Ints(nums)
    for i, x := range nums {
        if i % 2 == 0 {
            res += x
        } 
    }
    return res
}
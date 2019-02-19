package medium

// @author grekz
import "sort"
func minMoves2(nums []int) int {
    sort.Ints(nums)
    i, res, j := 0, 0, len(nums) - 1
    for i < j {
        res += nums[j] - nums[i]
        j-=1
        i+=1   
    }
    return res
}
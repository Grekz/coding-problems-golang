package medium

// @author grekz
import "strconv"
func summaryRanges(nums []int) []string {
    n, i, j := len(nums), 0, 0 
    res := make([]string, 0)
    for i < n {
        j = i
        c := strconv.Itoa(nums[i])
        for j + 1 < n && nums[j+1] - nums[j] == 1 {
            j++
        }
        if i == j {
            res = append(res, c)
        }else{
            res = append(res, c + "->" + strconv.Itoa(nums[j]))
        }
        i = j + 1
    }
    return res
}
package medium

// @author grekz
func findDuplicates(nums []int) []int {
    var res []int
    for i := 0; i < len(nums); i++ {
        cur := abs(nums[i]) - 1
        if nums[cur] < 0 {
            res = append(res, cur + 1)
        }else{
            nums[cur] = -nums[cur]        
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
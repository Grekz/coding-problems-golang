package medium

import "sort"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res := nums[0] + nums[1] + nums[2]
    le := len(nums) - 1
    for i:=0; i < le -1 ; i++ {
        cur := nums[i]
        if i > 0 && nums[i-1] == cur {
            continue
        }
        l := i + 1
        h := le 
        for l < h {
            tmp := cur + nums[l] + nums[h]
            if tmp == target {
                return tmp
            }
            if abs(tmp - target) < abs(res - target) {
                res = tmp
            }
            if tmp < target {
                l += 1                
            } else {
                h -= 1
            }
        }
    }
    return res
}
func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}
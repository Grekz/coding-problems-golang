package medium

import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    le := len(nums)
    end := le - 2
    res := [][]int{}
    for i, x := range nums {
        if x > 0 || i == end {
            break
        }
        if i > 0 && x == nums[i - 1] {
            continue
        }
        l := i + 1
        h := le - 1
        for l < h {
            tmp := x + nums[l] + nums[h]
            if tmp < 0 {
                l += 1
            }else if tmp > 0 {
                h -= 1
            }else{
                res = append(res, []int{x, nums[l], nums[h] })
                for l < h && nums[l] == nums[l+1] { l+=1 }
                for l < h && nums[h] == nums[h-1] { h-=1 }
                h-=1
                l+=1
            }
        }
    }
    return res
}
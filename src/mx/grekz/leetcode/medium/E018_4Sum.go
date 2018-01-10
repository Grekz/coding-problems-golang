package medium

import "sort"

func fourSum(nums []int, target int) [][]int {
    res := [][]int{}
    if ( len(nums) < 4 ) {return res}
    sort.Ints(nums)
    le := len(nums)
    ma := nums[le - 1]
    maxSum := ma * 3
    for i := 0; i < le - 3; i++ {
        x := nums[i]
        if ( x << 2 > target ){
            return res 
        }
        if ( i > 0 && x == nums[ i - 1 ] ){
            continue
        }
        if ( x + maxSum < target ){
            continue
        }
        for j := i+1; j < le - 2; j++ {
            if ( x + nums[j] * 3 > target ){
                break
            }
            if ( j > i +1 && nums[j] == nums[j-1] ){
                continue
            }
            cc := x + nums[j]
            l := j + 1
            h := le - 1
            for ( l < h ) {
                tmp := cc + nums[l] + nums[h]
                if ( tmp < target ){
                    l+=1
                }else if ( tmp > target ){
                    h-=1
                }else {
                    res = append(res, []int{x, nums[j], nums[l], nums[h]})
                    for ( l < h && nums[l] == nums[l+1] ){l+=1 }
                    for( l < h && nums[h] == nums[h-1] ){ h-=1 }
                    h-=1
                    l+=1
                }
            }
        }
    }
    return res
}
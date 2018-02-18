package easy

func majorityElement(nums []int) int {
    if len(nums) < 1 { return -1 }
    if len(nums) < 3 { return nums[0] }
    res := 0
    count := 0
    for _, x := range nums {
        if count == 0 { res = x }
        if res == x { 
            count++
        }else{
            count--
        }
    }
    return res
}
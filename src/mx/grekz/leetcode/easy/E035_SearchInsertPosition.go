package easy

func searchInsert(nums []int, target int) int {
    a := 0
    b := len(nums) - 1
    for a <= b {
        m := ( a + b ) / 2
        if ( nums[m] == target ) {return m}
        if ( nums[m] > target ){
            b = m - 1
        }else{
            a = m + 1
        }
    }
    return a
}
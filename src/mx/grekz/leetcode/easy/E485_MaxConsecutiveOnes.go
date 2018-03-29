package easy

// @author grekz
func findMaxConsecutiveOnes(nums []int) int {
    cur := 0
    ma := 0
    for _, x := range nums {
        if ( x == 1 ){
            cur+=1
        }else{
            cur = 0
        }
        if ( ma < cur ){
            ma = cur
        }
    }
    return ma
}
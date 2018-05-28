package medium

// @author grekz
func increasingTriplet(nums []int) bool {
    a, b := 2147483647, 2147483647
    for _, x := range nums {
        if x <= a {
            a = x
        }else if x <= b {
            b = x
        }else {
            return true
        }
    }
    return false
}
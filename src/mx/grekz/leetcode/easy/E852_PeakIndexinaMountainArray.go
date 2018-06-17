package easy

// @author grekz
func peakIndexInMountainArray(A []int) int {
    res, m := 0, -2147483648
    for i, x := range A {
        if x > m {
            res = i
            m = x
        }
    }
    return res
}
package medium

// @author grekz
func grayCode(n int) []int {
    res := make([]int, 1 << uint(n))
    for i := 0; i < len(res); i++ {
        res[i] = i ^ i >> 1
    }
    return res
}
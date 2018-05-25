package medium

// @author grekz
func uniquePaths(m int, n int) int {
    res := 1
    for i := 1; i < m; i++ {
        res = res * ( i + n - 1 ) / i
    }
    return res
}
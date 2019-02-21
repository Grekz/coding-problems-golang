package medium

// @author grekz
import "sort"
func smallestRangeII(A []int, K int) int {
    sort.Ints(A)
    n := len(A) - 1
    res := A[n] - A[0]
    for i := 0; i < n; i++  {
        a, b := A[i], A[i+1]
        h := max(A[n] - K, a + K)
        l := min(A[0] + K, b - K)
        res = min(res, h - l)
    }
    return res
}
func max( a,b int) int {
    if a > b {
        return a
    }
    return b
}
func min( a,b int) int {
    if a < b {
        return a
    }
    return b
}
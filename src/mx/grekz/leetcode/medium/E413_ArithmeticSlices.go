package medium

// @author grekz
func numberOfArithmeticSlices(A []int) int {
    tmp, res, n := 0, 0, len(A)
    for i := 2; i < n; i++ {
        if A[i] - A[i-1] == A[i-1] - A[i-2] {
            tmp++
            res += tmp
        }else{
            tmp = 0
        }
    }
    return res
}
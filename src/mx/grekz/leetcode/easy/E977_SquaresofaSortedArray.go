package easy

// @author grekz
func sortedSquares(A []int) []int {
    la := len(A)
    res := make([]int, la)
    for i := 0; i < la; i++ {
        A[i] = double(A[i])
    }
    for i, j, k := 0, la - 1, la - 1; i <= j; {
        if A[i] < A[j] {
            res[k] = A[j]
            j--
        }else{
            res[k] = A[i]
            i++
        }
        k--
    }
    return res
}
func double(x int) int {
    return x * x
}
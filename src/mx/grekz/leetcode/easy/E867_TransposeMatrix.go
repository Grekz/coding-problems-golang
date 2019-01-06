package easy

// @author grekz
func transpose(A [][]int) [][]int {
    res := make([][]int, len(A[0]))
    n := len(A)
    for i := range res {
        res[i] = make([]int, n)
    }
    for i, a := range A {
        for j, b := range a {
            res[j][i] = b
        }
    }
    return res
}
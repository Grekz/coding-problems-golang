package easy

// @author grekz
func isMonotonic(A []int) bool {
    a, b := true, true
    for i := 1; i < len(A); i++ {
        cur, prev := A[i], A[i-1]
        a = a && prev <= cur
        b = b && prev >= cur
    }
    return a || b
}
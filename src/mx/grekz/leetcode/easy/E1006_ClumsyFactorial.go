package easy

// @author grekz
func clumsy(N int) int {
    m := []int{1, 2, 2, -1, 0, 0, 3, 3}
    if N > 4 {
        return N + m[N % 4]    
    }
    return N + m[N + 3 ]
}
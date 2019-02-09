package medium

// @author grekz
func dailyTemperatures(T []int) []int {
    top, n := -1, len(T)
    stack, res := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        for top >= 0 && T[i] > T[stack[top]] {
            idx := stack[top]
            top--
            res[idx] = i - idx
        }
        top++
        stack[top] = i
    }
    return res
}
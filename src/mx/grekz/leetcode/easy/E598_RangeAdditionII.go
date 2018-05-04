package easy

// @author grekz
func maxCount(m int, n int, ops [][]int) int {
    row, col := m, n
    for _, x := range ops {
        row = Min(row, x[0])
        col = Min(col, x[1])
    }
    return row * col
}
func Min( a int, b int) int{
    if a < b {
        return a
    }
    return b
}
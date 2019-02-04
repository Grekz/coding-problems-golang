package hard

// @author grekz
func totalNQueens(n int) int {
    cols, d1, d2 := make([]bool, n), make([]bool, 2*n), make([]bool, n*2)
    return bt(cols, d1, d2, 0, n, 0)
}
func bt(cols, d1, d2 []bool, row, n, count int) int {
    if row == n {
        return count + 1
    }
    for i := 0; i < n; i++ {
        id1 := i - row + n
        id2 := i + row
        if !( cols[i] || d1[id1] || d2[id2] ) {
            cols[i] = true 
            d1[id1] = true 
            d2[id2] = true
            count = bt( cols, d1, d2, row + 1, n, count)
            cols[i] = false 
            d1[id1] = false 
            d2[id2] = false
        }
    }
    return count
}
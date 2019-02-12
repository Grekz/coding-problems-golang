package medium

// @author grekz
func matrixScore(A [][]int) int {
    rows, cols := len(A), len(A[0])
    var ans int =.0
    for c := 0; c < cols; c++ {
        var col int = 0
        for r := 0; r < rows; r++ {
            col += A[r][c] ^ A[r][0]
        }
        ans += max(col, rows - col) * ( 1 << uint( cols - c - 1 ) )
    }
    return ans
}

func max( a, b int) int {
    if a < b {
        return b
    }
    return a
}
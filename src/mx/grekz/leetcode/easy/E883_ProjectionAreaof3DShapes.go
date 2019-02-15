package easy

// @author grekz
func projectionArea(grid [][]int) int {
    res, n := 0, len(grid)
    for i := 0; i < n; i++ {
        row, col := 0, 0
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                res++
            }
            col = max(col, grid[j][i])
            row = max(row, grid[i][j])
        }
        res += col + row
    }
    return res
}

func max( a,b int) int {
    if a < b {
        return b
    }
    return a
}
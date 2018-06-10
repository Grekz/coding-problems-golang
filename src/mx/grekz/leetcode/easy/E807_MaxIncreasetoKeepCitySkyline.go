package easy

// @author grekz
func maxIncreaseKeepingSkyline(grid [][]int) int {
    n, m, res := len(grid), len(grid[0]), 0
    r, c := make([]int, n), make([]int, m)
    for i, x := range grid {
        for j, y := range x {
            if r[i] < y {
                r[i] = y
            }
            if c[j] < y {
                c[j] = y
            }
        }
    }
    for i, x := range grid {
        for j, y := range x {
            res += min(c[j], r[i]) - y
        }
    }
    return res
}
func min( a, b int) int {
    if a < b {return a}
    return b
}
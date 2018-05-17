package easy

// @author grekz
func maxAreaOfIsland(grid [][]int) int {
    ma := 0
    for i, x := range grid {
        for j, y := range x {
            if y == 1 {
                ma = max(ma, doit(grid,i,j))
            }
        }
    }
    return ma
}
func max( a, b int) int {
    if a < b {
        return b
    }
    return a
}
func doit(grid [][]int, x, y int) int {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
        return 0
    }
    grid[x][y] = 0
    return 1 + doit(grid,x-1,y) + doit(grid,x+1,y) + doit(grid,x,y-1) + doit(grid,x,y+1)
}
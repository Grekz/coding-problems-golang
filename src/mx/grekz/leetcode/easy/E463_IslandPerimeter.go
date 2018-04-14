package easy

// @author grekz
func islandPerimeter(grid [][]int) int {
    la, lb, isla, neig := len(grid), len(grid[0]), 0, 0
    for i:= 0; i < la; i++ {
        for j := 0; j < lb; j++ {
            if grid[i][j] == 1 {
                isla += 1
                if i < la - 1 && grid[i+1][j] == 1 {
                    neig += 1
                }
                if j < lb - 1 && grid[i][j+1] == 1 {
                    neig += 1
                }
            }
        }
    }
    return isla * 4 - neig * 2
}
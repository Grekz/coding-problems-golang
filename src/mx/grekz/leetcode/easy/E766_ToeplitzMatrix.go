package easy

// @author grekz
func isToeplitzMatrix(matrix [][]int) bool {
    m := len(matrix) - 1
    n := len(matrix[0]) - 1
    for i := 0; i < m; i++  {
        for j := 0 ; j < n; j++ {
            if matrix[i][j] != matrix[i+1][j+1] {
                return false
            }
        }
    }
    return true
}
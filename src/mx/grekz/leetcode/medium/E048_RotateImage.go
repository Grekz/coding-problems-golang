package medium

// @author grekz
func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            t := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = t
        }
    }
    n2 := n / 2
    for i := 0; i < n; i++ {
        for j := 0; j < n2; j++ {
            t := matrix[i][j]
            matrix[i][j] = matrix[i][n - j - 1]
            matrix[i][n - j - 1] = t
        }
    }
}
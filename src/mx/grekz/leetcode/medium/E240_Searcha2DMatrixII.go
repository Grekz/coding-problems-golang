package medium

// @author grekz
func searchMatrix(matrix [][]int, target int) bool {
    rs := len(matrix)
    if rs > 0 {
        r, c := 0, len(matrix[0]) - 1
        for c > -1 && r < rs {
            cur := matrix[r][c]
            if cur > target {
                c -= 1
            }else if cur < target {
                r += 1 
            }else{
                return true
            }
        }
    }
    return false
}
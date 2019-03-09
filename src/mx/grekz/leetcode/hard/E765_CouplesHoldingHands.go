package hard

// @author grekz
func minSwapsCouples(row []int) int {
    n, res := len(row), 0
    pos := make([]int, n, n)
    for i, x := range row {
        pos[x] = i
    }
    for i := 0; i < n; i += 2 {
        if pair := row[i] ^ 1; pair != row[i+1] {
            row[pos[pair]] = row[i+1]
            pos[row[i+1]] = pos[pair]
            res++
        }
    }
    return res
}
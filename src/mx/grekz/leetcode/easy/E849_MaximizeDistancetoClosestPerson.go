package easy

// @author grekz
func maxDistToClosest(seats []int) int {
    i, res, n := 0, 0, len(seats)
    for j, cur := range seats {
        if cur == 1 {
            if i == 0 {
                res = max(res, j - i)
            }else{
                res = max(res, (j - i + 1) / 2)
            }
            i = j + 1
        }
    }
    return max(res, n - i)
}
func max( a,b int) int {
    if a > b {
        return a
    }
    return b
}
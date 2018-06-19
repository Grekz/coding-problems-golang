package hard

// @author grekz
func longestConsecutive(nums []int) int {
    res, s := 0, make(map[int]bool)
    for _, x := range nums { s[x] = true }
    for _, x := range nums {
        if !contains(x - 1, s) {
            y := x + 1
            for contains(y, s) {
                y++
            }
            res = max(res, y - x )
        }
    }
    return res
}
func contains( x int, m map[int]bool) bool{
    if _, x := m[x]; x{
        return true
    }
    return false
}
func max ( a,b int) int {
    if a > b { return a }
    return b
}
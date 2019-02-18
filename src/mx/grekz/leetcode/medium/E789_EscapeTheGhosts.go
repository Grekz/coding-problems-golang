package medium

// @author grekz
func escapeGhosts(ghosts [][]int, target []int) bool {
    t0, t1 := target[0], target[1]
    dest := abs(t0) + abs(t1)
    for _, g := range ghosts {
        cur := abs(g[0]-t0) + abs(g[1]-t1)
        if cur <= dest {
            return false
        }
    }
    return true
}
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
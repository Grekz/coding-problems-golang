package medium

// @author grekz
func isSubsequence(s string, t string) bool {
    idx := -1
    for _, c := range(s) {
        idx = findIdx(t, byte(c), idx + 1)
        if idx == -1 {
            return false
        }
    }
    return true
}
func findIdx( t string, c byte, idx int) int {
    for i := idx; i < len(t); i++ {
        if t[i] == c {
            return i
        }
    }
    return -1
}
package easy

// @author grekz
func countBinarySubstrings(s string) int {
    res, p, c, n := 0, 0, 1, len(s)
    for i := 1; i < n; i++ {
        if s[i] != s[i-1] {
            res += min(p,c)
            p = c
            c = 1
        }else{
            c += 1
        }
    }
    return res + min(p,c)
}
func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
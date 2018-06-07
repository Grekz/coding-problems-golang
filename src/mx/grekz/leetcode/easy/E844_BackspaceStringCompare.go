package easy

// @author grekz
func backspaceCompare(S string, T string) bool {
    return doit(S) == doit(T)
}
func doit(s string) string {
    n, c, res, hash := len(s), 0, "", byte('#')
    for i := n - 1; i > -1; i-- {
        cur := s[i]
        if cur == hash {
            c += 1
        } else if c > 0 {
            c -= 1
        } else {
            res += string(cur)
        }
    }
    return res
}
package easy

// @author grekz
func removeOuterParentheses(S string) string {
    c, res := 0, ""
    for _, x := range S {
        if x == '(' {
            if c > 0 {
                res += string(x)
            }
            c += 1
        }
        if x == ')' {
            if c > 1 {
                res += string(x)
            }
            c -= 1
        }
    }
    return res
}
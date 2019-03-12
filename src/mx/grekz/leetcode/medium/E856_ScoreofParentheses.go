package medium

// @author grekz
func scoreOfParentheses(S string) int {
    bal, res, prev := 0, 0, rune('0')
    for _, c := range S {
        if  c == '(' {
            bal++
        }else {
            bal--
            if prev == '(' {
                res += 1 << uint(bal)
            }
        }
        prev = c
    }
    return res
}
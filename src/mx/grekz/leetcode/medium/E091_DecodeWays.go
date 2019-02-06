package medium

// @author grekz
func numDecodings(s string) int {
    if len(s) < 1 || s[0] == byte('0') {
        return 0
    }
    res := make([]int, len(s) + 1)
    res[1], res[0] = 1, 1
    for i := 2; i <= len(s); i++ {
        a, b := int(s[i-2]), int(s[i-1])
        if b != 48 {
            res[i] = res[i-1]
        }
        if a == 49 || ( a == 50 && b < 55 ) {
            res[i] += res[i-2]
        }
    }
    return res[len(res)-1]
}
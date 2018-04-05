package easy

// @author grekz
func longestPalindrome(s string) int {
    tmp := make([]bool, 58,58)
    res := 0
    for _, x := range s {
        t := int(x) - 65
        if tmp[t] {
            res += 2
        }
        tmp[t] = !tmp[t]
    }
    if res != len(s) {
        return 1 + res
    }
    return res
}
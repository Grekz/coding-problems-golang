package easy

func titleToNumber(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        res = res * 26 - 64 + int(s[i])
    }
    return res
}
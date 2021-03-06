package easy
import "strings"
func toHex(num int) string {
    if num == 0 { return "0" }
    res := ""
    cc := []rune{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'}
    for i := 0; i < 8; i++ {
        res = string(cc[num & 15]) + res
        num >>=4
    }
    return strings.TrimLeft(res, "0")
}
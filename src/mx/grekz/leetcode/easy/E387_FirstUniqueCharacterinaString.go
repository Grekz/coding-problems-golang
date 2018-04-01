package easy

// @author grekz
import "strings"
func firstUniqChar(s string) int {
    letters := "abcdefghijklmnopqrstuvwxyz"
    res := len(s)
    for _, c := range letters {
        lInd := strings.Index(s, string(c))
        if lInd != -1 && lInd == strings.LastIndex(s, string(c)) && res > lInd {
            res = lInd
        }
    }
    if res == len(s) {
        return -1
    }
    return res
}
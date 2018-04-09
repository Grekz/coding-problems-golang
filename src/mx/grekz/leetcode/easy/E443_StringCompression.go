package easy

// @author grekz
import "strconv"
func compress(chars []byte) int {
    ind, ans := 0, 0
    n := len(chars)
    for ind < n {
        c := chars[ind]
        cc := 0
        for ind < n && c == chars[ind] {
            ind += 1
            cc += 1
        }
        chars[ans] = c
        ans += 1
        if cc > 1 {
            for _, x := range strconv.Itoa(cc) {
                chars[ans] = byte(x)
                ans += 1
            }
        }
    }
    return ans
}
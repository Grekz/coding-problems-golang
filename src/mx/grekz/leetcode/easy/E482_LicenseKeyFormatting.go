package easy

// @author grekz
import "strings"
func licenseKeyFormatting(S string, K int) string {
    c := strings.Replace(strings.ToUpper(S), "-", "", -1)
    le := len(c)
    res := ""
    if le < 1 {
        return res
    }
    s1 := le % K
    if s1 == 0 {
        s1 = K
    }
    res = c[:s1]
    for le > s1 {
        tmp := s1 + K
        res += "-" + c[s1:tmp]
        s1 = s1+K
    }
    return res
}
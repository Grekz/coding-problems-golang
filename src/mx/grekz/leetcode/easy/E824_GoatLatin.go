package easy

// @author grekz
import "strings"
func toGoatLatin(S string) string {
    vowels := "aeiouAEIOU"
    res := ""
    i := 1
    for _, c := range strings.Fields(S) {
        curChar := string(c[0])
        res += " "
        if strings.ContainsAny(vowels, curChar) {
            res += c
        }else{
            res += c[1:len(c)] + curChar
        }
        res += "ma"
        for j := 0; j < i; j++ {
            res += "a"
        }
        i+=1
    }
    return res[1:]
}
package medium

// @author grekz
func findAndReplacePattern(words []string, pattern string) []string {
    res := make([]string, 0)
    for _, x := range words {
        if check(x, pattern) {
            res = append(res, x)
        }
    }
    return res
}

func check(a, b string) bool {
    dict, i := make(map[rune]rune, 0), 0
    for _, x := range a {
        bChar := rune(b[i])
        i++
        if _, ok := dict[x]; !ok {
            dict[x] = bChar
        }
        if dict[x] != bChar {
            return false
        }
    }
    ch := make([]bool, 26, 26)
    b97 := rune(97)
    for _, x := range getValues(dict) {
        if ch[x - b97] {
            return false
        }
        ch[x - b97] = true
    }
    return true
}

func getValues( m map[rune]rune ) []rune {
    v := make([]rune, 0, len(m))
    for  _, value := range m {
       v = append(v, value)
    }
    return v
}
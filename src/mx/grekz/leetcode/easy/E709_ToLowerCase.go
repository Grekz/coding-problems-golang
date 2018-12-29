package easy

// @author grekz
func toLowerCase(str string) string {
    res := ""
    for _, c := range str { 
        if 65 <= c && c <= 90 { 
            res += string(rune(c + 32))
        }else{
            res += string(c)
        }
    }
    return res
}
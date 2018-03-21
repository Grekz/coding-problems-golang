package easy

// @author grekz

import "strings"
func wordPattern(pattern string, str string) bool {
    words := strings.Fields(str)
    if (len(words) != len(pattern)) {return false}
    dic := map[byte]string{}
    for i, x := range words{
        char := pattern[i]
        if e, ok := dic[char]; ok{
            if ( e != x ) {
                return false
            }
        }else if containsValue(dic,x){
            return false
        }else{
            dic[char] = words[i]
        }
    }
    return true
}
func containsValue(m map[byte]string, v string) bool {
    for _, x := range m {
        if x == v {
            return true
        }
    }
    return false
}
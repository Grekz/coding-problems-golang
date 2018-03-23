package easy

// @author grekz
func reverseVowels(s string) string {
    r := []rune(s)
    for i, j := 0, len(s) - 1; i < j; {
        for i < j && check(r[i]) { i+=1 }
        for i < j && check(r[j]) { j-=1 }
        if i < j {
            r[i], r[j] = r[j], r[i]
        }
        i+=1
        j-=1
    }
    return string(r)
}
func check(r rune) bool{
    switch r {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return false
    }
    return true
}
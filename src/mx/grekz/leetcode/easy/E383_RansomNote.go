package easy

// @author grekz
func canConstruct(ransomNote string, magazine string) bool {
    mag := map[rune]int{}
    for _, x := range magazine {
        if _, ok := mag[x]; ok {
            mag[x] += 1
        }else{
            mag[x] = 1
        }
    }
    for _, x := range ransomNote {
        if e, ok := mag[x]; !ok || e == 0 {
            return false
        }
        mag[x] -= 1
    }
    return true
}
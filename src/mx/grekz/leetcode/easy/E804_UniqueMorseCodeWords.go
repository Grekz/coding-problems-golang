package easy

// @author grekz
func uniqueMorseRepresentations(words []string) int {
    cc := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    s := make(map[string]bool)
    for _, w := range words {
        x := ""
        for _, c := range w {
            x += cc[int(c) - 97]
        }
        s[x] = true
    }
    return len(s)
}
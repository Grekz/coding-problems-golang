package easy

import "strings"

func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    a := len(s)
    b := strings.LastIndex(s," ")
    return a - b - 1
}
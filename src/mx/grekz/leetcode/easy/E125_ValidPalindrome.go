package easy

import (
    "unicode"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        for i < j && !(unicode.IsDigit(rune(s[i])) || unicode.IsLetter(rune(s[i]))) { i+=1 }
        for i < j && !(unicode.IsDigit(rune(s[j])) || unicode.IsLetter(rune(s[j]))) { j-=1 }
        if ( s[i] != s[j] ) {
            return false
        }
    }
    return true
}
package medium

import (
    "strings"
)
func letterCombinations(digits string) []string {
    if len(digits) < 1 || strings.ContainsAny(digits, "0") || strings.ContainsAny(digits, "1") {
        return []string{}
    }
    val := map[rune]string{
        '2' : "abc",
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
    }
    res := make([]string, 0)
    res = append(res,"")
    for _, x := range digits {
        tmp := make([]string,0)
        for _, y := range val[x] {
            for _, z := range res {
                tmp = append(tmp, string(z) + string(y))
            }
        }
        res = tmp
    }
    return res
}
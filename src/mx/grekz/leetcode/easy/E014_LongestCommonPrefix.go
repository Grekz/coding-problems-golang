package easy

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    res := strs[0]
    strs = strs[1:]
    for _, x := range strs {
        for !strings.HasPrefix( x, res ) {
            res = res[:len(res)-1]
        }
    }
    return res
}
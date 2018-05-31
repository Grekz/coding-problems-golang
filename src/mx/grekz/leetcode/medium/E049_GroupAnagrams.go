package medium

// @author grekz
import (
    "sort"
    "strings"
)
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for _, x := range strs {
        key := sorted(x)
        c, ok := m[key]
        if !ok {
            c = make([]string, 0)
        }
        m[key] = append(c, x)
    }
    res := make([][]string, 0)
    for _, v := range m {
        res = append(res, v)
    }
    return res
}

func sorted(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
package medium

// @author grekz
func customSortString(S string, T string) string {
    tp := make([]int, 26, 26)
    res := ""
    for _, x := range T {
        tp[x-97]++
    }
    for _, x := range S {
        for tp[x-97] > 0 {
            res += string(x)
            tp[x-97]--
        }
    }
    for _, x := range T {
        for tp[x-97] > 0 {
            res += string(x)
            tp[x-97]--
        }
    }
    return res
}
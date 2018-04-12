package easy

// @author grekz
import "sort"
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    gi, si := 0,0
    gn, sn := len(g), len(s)
    for gi < gn && si < sn {
        if g[gi] <= s[si] {
            gi += 1
        }
        si += 1
    }
    return gi
}
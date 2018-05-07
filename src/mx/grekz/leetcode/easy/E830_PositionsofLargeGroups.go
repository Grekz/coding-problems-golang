package easy

// @author grekz
func largeGroupPositions(S string) [][]int {
    res := make([][]int, 0)
    ps := 0
    rps := rune(S[ps])
    S += "|"
    for i, x := range S {
        if x != rps {
            if i - ps > 2 {
                res = append(res, []int{ps, i - 1})
            }
            ps = i
            rps = rune(S[ps])
        }
    }
    return res
}
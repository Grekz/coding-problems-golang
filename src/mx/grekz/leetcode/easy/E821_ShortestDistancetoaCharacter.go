package easy

// @author grekz
func shortestToChar(S string, C byte) []int {
    n := len(S)
    res := make([]int, n)
    pos := -n
    runeC := rune(C)
    for i, x := range S {
        if x == runeC {
            // res[i] = 0
            pos = i
        }else{
            res[i] = i - pos
        }
    }
    // fmt.Println(res)
    for i := n -1; i >= 0; i-- {
        cur := S[i]
        if cur == C {
            pos = i
        }else{
            res[i] = Min(res[i], Abs(i - pos))
        }
    }
    return res
}
func Abs(a int) int{
    if a < 0 {
        return -a
    }
    return a
}
func Min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}
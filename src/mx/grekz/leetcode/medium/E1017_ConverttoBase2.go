package medium

// @author grekz
func baseNeg2(N int) string {
    res := string("")
    m := []string{"0", "1"}
    for N != 0 {
        res = m[N & 1] + res
        N = -( N >> 1 )
    }
    if len(res) > 0 {
        return res
    }
    return "0"
}
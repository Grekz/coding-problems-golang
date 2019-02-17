package easy

// @author grekz
func binaryGap(N int) int {
    res, l := 0, -1
    for i := 0; i < 32; i++ {
        if ( ( N >> uint(i) ) & 1 ) > 0 {
            if l >= 0 {
                res = max(res, i - l)
            }
            l = i
        }
    }
    return res
}
func max(a, b int ) int {
    if a > b { 
        return a
    }
    return b
}
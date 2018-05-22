package hard

// @author grekz
func largestRectangleArea(heights []int) int {
    s := make([]int, 0)
    s = append(s,0)
    heights = append(heights, 0)
    res := 0
    for i, x := range heights {
        for len(s) > 0 && x < heights[s[len(s)-1]] {
            h := heights[s[len(s)-1]]
            s = s[:len(s)-1]
            w := i
            if len(s) > 0 {
                w = w - s[len(s)-1] - 1
            }
            res = max(res, h * w)
        }
        s = append(s, i)
    }
    return res
}
func max( a, b int ) int {
    if a > b {
        return a
    }
    return b
}
package easy

// @author grekz
func largestTriangleArea(points [][]int) float64 {
    res := 0.0
    for _, i := range points {
        for _, j := range points {
            for _, k := range points {
                res = Max(res, 0.5 * Abs(i[0] * j[1] + j[0] * k[1] + k[0] * i[1]- j[0] * i[1] - k[0] * j[1] - i[0] * k[1]))
            }
        }
    }
    return res
}
func Max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
func Abs(a int) float64 {
    if a < 0 {
        a = -a
    }
    return float64(a)
}
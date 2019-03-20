package medium

// @author grekz
func maxProfit(prices []int, fee int) int {
    a, b, c := 0, -2147483648, 0
    for _, x := range prices {
        c = a
        a = max( a, b + x )
        b = max( b, c - fee - x)
    }
    return a
}
func max( a, b int ) int {
    if a < b {
        return b
    }
    return a
}
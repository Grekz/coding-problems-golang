package medium

// @author grekz
func maxProfit(prices []int) int {
    a,b,c := 0,0, -2147483648
    for _, price := range prices { 
        d := b
        b = max(b, c + price)
        c = max(c, a - price)
        a = d    
    }
    return b
}
func max ( a, b int) int {
    if a > b {
        return a
    }
    return b
}
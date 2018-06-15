package hard

// @author grekz
func maxProfit(prices []int) int {
    profitA, profitB, buyA, buyB := 0,0,-2147483648, -2147483648
    for i, cur := range prices {
        if i > 2 {
            profitB = max(profitB, buyB + cur)
        }
        if i > 1 {
            buyB = max(buyB, profitA - cur)
        }
        profitA = max(profitA, buyA + cur)
        buyA = max(buyA, -cur)
    }
    return max(profitB, profitA)
}
func max( a,b int) int {
    if a > b {
        return a
    }
    return b
}
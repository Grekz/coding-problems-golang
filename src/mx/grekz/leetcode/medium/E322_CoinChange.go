package medium

// @author grekz
func coinChange(coins []int, amount int) int {
    m := amount + 1
    dp := make([]int, m, m)
    dp[0] = 0
    for i := 1; i < m; i++ {
        dp[i] = m
        for _, c := range coins {
            if c <= i {
                dp[i] = min(dp[i], dp[i - c] + 1)
            }
        }
    }
    
    if dp[amount] <= amount {
        return dp[amount]
    }
    return -1
}
func min( a, b int) int {
    if a < b {
        return a
    }
    return b
}
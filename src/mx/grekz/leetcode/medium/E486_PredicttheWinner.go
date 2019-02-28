package medium

// @author grekz
func PredictTheWinner(nums []int) bool {
    n := len(nums)
    dp := make([]int, n, n)
    for i, _ := range nums {
        dp[i] = 0
    }
    for i := n; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            a := nums[i] - dp[j]
            b := nums[j] - dp[j - 1]
            dp[j] = max(a, b)
        }
    }
    return dp[n-1] >= 0
}
func max( a, b int) int {
    if a < b {
        return b
    }
    return a
}
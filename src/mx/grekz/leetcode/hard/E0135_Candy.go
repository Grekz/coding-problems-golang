package hard

// @author grekz
func candy(ratings []int) int {
    n := len(ratings)
    tmp := make([]int, n, n)
    for i := 0; i < n; i++ {
        tmp[i] = 1
    }
    for i := 1; i < n; i++ {
        if ratings[i] > ratings[i-1] {
            tmp[i] = tmp[i-1] + 1
        }
    }
    res := tmp[n-1]
    for i := n-2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            tmp[i] = max(tmp[i+1] + 1, tmp[i])
        }
        res += tmp[i]
    }
    return res
}
func max( a, b int ) int {
    if a > b {
        return a
    }
    return b
}
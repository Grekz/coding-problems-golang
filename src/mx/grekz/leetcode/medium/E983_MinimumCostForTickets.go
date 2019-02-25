package medium

// @author grekz
func mincostTickets(days []int, costs []int) int {
    mint, n := 2147483647, days[len(days) - 1]
    cd := make([]int, n + 1)
    for _, x := range(days) {
        cd[x] = mint
    }
    for i := 1; i <= n; i++ {
        if cd[i] == mint {
            mc := cd[i-1] + costs[0]
            mc = min(mc, cd[goz(i, 7)] + costs[1])
            mc = min(mc, cd[goz(i, 30)] + costs[2])
            cd[i] = mc
        }else{
            cd[i] = cd[i-1]
        }
    } 
    return cd[n]
}
func min( a, b int) int {
    if a < b {
        return a
    }
    return b
}
func goz( a, b int) int {
    if a < b {
        return 0
    }
    return a - b
}
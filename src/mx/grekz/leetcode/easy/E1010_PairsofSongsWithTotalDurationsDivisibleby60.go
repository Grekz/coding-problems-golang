package easy

// @author grekz
func numPairsDivisibleBy60(time []int) int {
    res, cnt := 0, make([]int, 60)
    for _, x := range time {
        d := (60 - x % 60) % 60
        res += cnt[d]
        cnt[x % 60] += 1
    }
    return res
}
package easy

// @author grekz
func distributeCandies(candies []int) int {
    n := len(candies) / 2
    m := make(map[int]bool)
    for _, x := range candies {
        m[x] = true
    }
    mm := len(m)
    if n < mm {
        return n
    }
    return mm
}
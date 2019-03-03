package medium

// @author grekz
func kthGrammar(N int, K int) int {
    n, k := 0, K - 1
    for k > 0 {
        k &= k - 1
        n += 1
    }
    return n & 1
}
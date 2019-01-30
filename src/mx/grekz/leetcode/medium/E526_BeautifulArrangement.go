package medium

// @author grekz
func countArrangement(N int) int {
    return permutate(N, 1, 0, make([]bool, N + 1))
}
func permutate(n, pos, count int, v []bool) int {
    if pos > n {
        return count + 1
    }
    for i := 1; i <= n; i++ {
        if !v[i] && (i % pos == 0 || pos % i == 0) {
            v[i] = true
            count = permutate(n, pos + 1, count, v)
            v[i] = false
        }
    }
    return count
}
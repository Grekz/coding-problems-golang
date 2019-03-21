package easy

// @author grekz
func findJudge(N int, trust [][]int) int {
    c := make([]int, N + 1)
    for _, x := range trust {
        c[x[0]] -= 1
        c[x[1]] += 1
    }
    for i := 1; i <= N; i++  {
        if c[i] == N - 1 {
            return i
        }
    }
    return -1
}
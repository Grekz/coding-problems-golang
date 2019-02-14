package medium

// @author grekz
func beautifulArray(N int) []int {
    res := []int{1}
    for len(res) < N {
        tmp := make([]int, 0)
        for _, i := range res {
            if i * 2 - 1 <= N {
                tmp = append(tmp, i * 2 - 1)
            }
        }
        for _, i := range res {
            if i * 2 <= N {
                tmp = append(tmp, i * 2)
            }
        }
        res = tmp
    }
    return res
}
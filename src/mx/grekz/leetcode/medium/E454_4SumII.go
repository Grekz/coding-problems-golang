package medium

// @author grekz
func fourSumCount(A []int, B []int, C []int, D []int) int {
    m := make(map[int]int)
    res := 0
    for _, a := range A {
        for _, b := range B {
            m[a+b] += 1
        }
    }
    for _, a := range C {
        for _, b := range D {
            res += m[-(a+b)]
        }
    }
    return res
}
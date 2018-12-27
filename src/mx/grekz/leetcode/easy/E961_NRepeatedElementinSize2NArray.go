package easy

// @author grekz
func repeatedNTimes(A []int) int {
    c := make(map[int]bool)
    for _, x := range A {
        if c[x] {
            return x
        }
        c[x] = true
    }
    return -1
}
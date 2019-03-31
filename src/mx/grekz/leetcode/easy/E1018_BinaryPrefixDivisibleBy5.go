package easy

// @author grekz
func prefixesDivBy5(A []int) []bool {
    tmp, res := 0, make([]bool, 0)
    for _, x := range A {
        tmp = ( ( tmp << 1 ) + x ) % 5
        res = append(res, tmp == 0)
    }
    return res
}
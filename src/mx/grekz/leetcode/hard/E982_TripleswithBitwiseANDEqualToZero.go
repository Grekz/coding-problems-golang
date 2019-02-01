package hard

// @author grekz
func countTriplets(A []int) int {
    res, tmp  := 0, make(map[int]int)
    for _, a := range A {
        for _, b := range A {
            if _, ok := tmp[a&b]; ok {
                tmp[a&b]++
            }else{
                tmp[a&b] = 1
            }
        }
    }
    for k, v := range tmp {
        for _, b := range A {
            if k & b == 0 {
                res += v
            }
        }
    }
    return res
}
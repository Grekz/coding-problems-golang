package easy

// @author grekz
func sortArrayByParityII(a []int) []int {
    i, j, res := 0, 1, make([]int, len(a))
    for _, x := range a {
        if ( x & 1) == 0 {
            res[i] = x
            i += 2
        }else{
            res[j] = x
            j += 2
        }
    }
    return res
}
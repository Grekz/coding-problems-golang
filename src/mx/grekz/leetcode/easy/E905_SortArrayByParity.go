package easy

// @author grekz
func sortArrayByParity(a []int) []int {
    i, j := 0, len(a) - 1
    for i < j {
        if cur := a[i]; cur % 2 == 1 {
            a[i] = a[j]
            a[j] = cur
            j -= 1
        }else{
            i += 1
        }
    }
    return a
}
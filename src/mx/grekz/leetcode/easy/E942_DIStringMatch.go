package easy

// @author grekz
func diStringMatch(S string) []int {
    n := len(S)
    right, left := len(S), 0
    inc := rune('I')
    res := make([]int, n + 1)
    for i, x := range S {
        cur := left
        if x == inc {
            left += 1
        }else{
            cur = right
            right -= 1
        }
        res[i] = cur 
    }
    res[n] = right
    return res
}
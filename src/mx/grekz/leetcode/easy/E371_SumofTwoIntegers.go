package easy

// @author grekz
func getSum(a int, b int) int {
    tMAX := 0x7FFFFFFF
    mask := 0xFFFFFFFF
    for b != 0 {
        tmp := a
        a =  ( a ^ b ) & mask
        b = ((tmp & b) << 1) & mask
    }
    if a <= tMAX {
        return a
    }
    return  ^(a ^ mask)
}
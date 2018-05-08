package easy

// @author grekz
func isOneBitCharacter(bits []int) bool {
    n := len(bits) - 1
    if n < 1 {
        return true
    }
    for i := n - 1; i >= 0; i-- {
        if bits[i] == 0 {
            return ( n - i ) % 2 != 0
        }
    }
    return ( n - 1 ) % 2 != 0
}
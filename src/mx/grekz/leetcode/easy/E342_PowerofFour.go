package easy

// @author grekz
func isPowerOfFour(num int) bool {
    if ( num < 1 ) { return false }
    for num % 4 == 0 {
        num /= 4
    }
    return num == 1
}
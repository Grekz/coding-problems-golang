package easy

// @author grekz
func addDigits(num int) int {
    if(num < 10) {return num}
    return 1 + ( num - 1 ) % 9
}
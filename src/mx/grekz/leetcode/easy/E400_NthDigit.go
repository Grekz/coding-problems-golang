package easy

// @author grekz
import "strconv"
func findNthDigit(n int) int {
    if ( n < 10 ) { 
        return n
    }
    n -= 1
    digitInNumber := 1
    rangeNumber := 1
    for n / 9 / rangeNumber / digitInNumber >= 1 {
        n -= 9 * rangeNumber * digitInNumber
        digitInNumber+=1
        rangeNumber *= 10
    }
    res := strconv.Itoa(rangeNumber + n/digitInNumber)
    return int(res[n%digitInNumber]) - '0'
}
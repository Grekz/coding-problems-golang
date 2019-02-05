package medium

// @author grekz
func countNumbersWithUniqueDigits(n int) int {
    if ( n == 0 ) { return 1 }
    if ( n == 1 ) { return 10 }
    if ( n == 2 ) { return 91 }
    if ( n > 10 ) { return 0 }
    res, t := 10, 9
    for i := 1; i < n; i++ {
        t *= 10 - i
        res += t
    }
    return res
}
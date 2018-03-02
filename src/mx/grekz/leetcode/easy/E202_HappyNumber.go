package easy

func isHappy(n int) bool {
    if ( n < 1 || n == 4 ) {
        return false
    }
    if ( n == 1 || n == 7 || n == 10 || n == 13 || n == 19 || n == 23 ) {
        return true
    }
    res := 0
    for n > 0 {
        res += n%10 * (n%10)
        n /= 10
    }
    return isHappy(res)
}
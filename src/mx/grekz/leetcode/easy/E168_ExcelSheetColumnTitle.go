package easy

func convertToTitle(n int) string {
    if ( n == 0 ) { return "" }
    res := ""
    for n > 0 {
        n -= 1
        res = string(65+n%26) + res
        n /= 26
    }
    return res
}
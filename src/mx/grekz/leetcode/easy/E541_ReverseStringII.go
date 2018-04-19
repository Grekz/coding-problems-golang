package easy

// @author grekz
func reverseStr(s string, k int) string {
    cArr := []byte(s)
    n := len(cArr)
    for i := 0; i < n; i += 2 * k {
        ii := i
        j := i + k - 1
        if j > n -1 {
            j = n - 1
        }
        for ii < j {
            tmp := cArr[ii]
            cArr[ii] = cArr[j]
            ii += 1
            cArr[j] = tmp
            j -= 1
        }
    }
    return string(cArr)
}
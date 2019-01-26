package medium

// @author grekz
import (
    "strings"
    "strconv"
)
func compareVersion(version1 string, version2 string) int {
    a, b := strings.Split(version1, "."), strings.Split(version2, ".")
    al, bl := len(a), len(b)
    n := max(al, bl)
    for i := 0 ; i < n; i++ {
        c := getInt(i, al, a) - getInt(i, bl, b)
        if c < 0 {
            return -1
        }else if c > 0 {
            return 1
        }
    }
    return 0
}
func getInt( i, l int, arr []string ) int {
    if i < l {
        res, _ := strconv.Atoi(arr[i])
        return res
    }
    return 0
}
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
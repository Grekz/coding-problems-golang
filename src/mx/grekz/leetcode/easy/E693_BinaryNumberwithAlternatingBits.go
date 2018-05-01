package easy

// @author grekz
import (
    "strconv"
    "strings"
)
func hasAlternatingBits(n int) bool {
    b := strconv.FormatInt(int64(n), 2)
    return !strings.Contains(b, "00") && !strings.Contains(b, "11")
}
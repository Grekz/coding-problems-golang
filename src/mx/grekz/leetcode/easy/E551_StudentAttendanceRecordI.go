package easy

// @author grekz
import "strings"
func checkRecord(s string) bool {
    return !strings.Contains(s, "LLL") && strings.Index(s, "A") == strings.LastIndex(s, "A") 
}
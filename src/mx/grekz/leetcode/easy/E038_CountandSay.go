package easy

// @author grekz
import "strconv" 
func countAndSay(n int) string {
    if ( n < 1 ) { return "0" }
    if ( n == 1 ) { return "1" }
    if ( n == 2 ) { return "11" }
    if ( n == 3 ) { return "21" }
    arr := countAndSay( n - 1 )
    cur, count := 'a', 0
    res := ""
    for _, x := range arr {
        if ( cur != x ) {
            if (count > 0) {
                res += strconv.Itoa(count) + "" + string(cur)
            }
            count = 1
            cur = x
        }else{
            count+=1
        }
    }
    res += strconv.Itoa(count) + "" + string(cur)
    return res
}
package medium

// @author grekz
func minAddToMakeValid(S string) int {
    res, bal := 0, 0
    for _, x := range(S) {
        if x == '(' {
            bal += 1
        } else {
            bal -= 1  
        }
        if bal == -1 {
            bal++
            res++
        }
    }
    return bal + res
}
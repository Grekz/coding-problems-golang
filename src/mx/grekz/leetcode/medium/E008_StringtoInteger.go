package medium

import (
    "strconv"
    "strings"
)
func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if len(str) < 1 {
        return 0
    }
    sign, res, z, n := byte('+'), "", '0', '9'
    if str[0] == '-' || str[0] == '+' {
        sign = str[0]
        str = str[1:]
    }
    for _, x := range str {
        if x >= z && x <= n {
            res += string(x)            
        }else{
            break
        }
    }
    if res != "" {
        resInt, _ := strconv.ParseInt( string(sign) + res, 10,0 )
        if  resInt > 2147483647 { return 2147483647 }
        if  resInt < -2147483648 { return -2147483648 }
        return int(resInt)
    }                
    return 0
}
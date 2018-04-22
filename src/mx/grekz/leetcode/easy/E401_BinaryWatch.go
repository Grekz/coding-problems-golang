package easy

// @author grekz
import (
    "math/bits"
    "fmt"
)
func readBinaryWatch(num int) []string {
    res:=[]string{}
    for i := uint(0); i < 12; i++ {
    	for j := uint(0); j < 60; j++ {
            if ( bits.OnesCount(i) + bits.OnesCount(j) == num ) {
                res = append(res, fmt.Sprintf("%d:%02d", i, j))
            }
        }
    }
    return res
}
package easy

// @author grekz
import "math/bits"
func countPrimeSetBits(L int, R int) int {
    res := 0
    tmp := uint(665772)
    fmt.Println(tmp)
    for ; L <= R; L++ {
        cur := (tmp >> uint(bits.OnesCount(uint(L))) & 1)
        res += int(cur)
    } 
    return res
}
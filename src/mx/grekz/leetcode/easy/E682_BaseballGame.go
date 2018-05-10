package easy

// @author grekz
import "strconv"
func calPoints(ops []string) int {
    res, val := 0, 0
    var rnd []int
    for _, x := range ops {
        l := len(rnd) - 1
        if x == "C" {
            val = -rnd[l]
            rnd = rnd[:l]
        }else{
            if x == "+" {
                val = rnd[l] + rnd[l-1]
            }else if x == "D" {
                val = 2 * rnd[l]
            }else{
                val, _ = strconv.Atoi(x)
            }
            rnd = append(rnd, val)
        }
        res += val
    }
    return res
}
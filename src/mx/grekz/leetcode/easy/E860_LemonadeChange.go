package easy

// @author grekz
func lemonadeChange(bills []int) bool {
    f, t := 0, 0
    for _, e := range bills {
        if e == 5 { 
            f += 1 
        }else if e == 10 {
            t += 1
            f -= 1
        }else if t > 0 {
            t -= 1
            f -= 1
        }else{
            f -= 3
        }
        if f < 0 {
            return false
        }
    }
    return true
}
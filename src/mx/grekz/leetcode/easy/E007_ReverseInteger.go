package easy

func reverse(x int) int {
    s, r := 1, 0
    if x < 0 {
        s = -1
    }
    x *= s
    
    for x > 0 {
        r = r * 10 + x % 10
        x /= 10
    }
    if r > 2147483647 {
        return 0
    }
    return r * s
}
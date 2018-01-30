package easy

func climbStairs(n int) int {
    a := math.Sqrt(float64(5))
    b := math.Pow(float64( float64(1 + a) / float64(2)),float64( n + 1))
    c := math.Pow(float64( float64(1 - a) / float64(2)), float64( n + 1))
    res := 1.0/float64(a) * float64( float64(b) - float64(c) )
    return int( res + 0.5 )
}

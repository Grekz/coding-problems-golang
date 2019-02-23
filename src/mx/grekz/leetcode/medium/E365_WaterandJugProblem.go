package medium

// @author grekz
func canMeasureWater(x int, y int, z int) bool {
    if x + y < z { return false }
    if x == z || y == z || x + y == z { return true }
    return z % gcd(y, x) == 0
}
func gcd( a,b int) int {
    if b == 0 {
        return a
    }
    return gcd( b, a % b)
}
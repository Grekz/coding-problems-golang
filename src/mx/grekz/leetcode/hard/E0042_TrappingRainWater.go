package hard

// @author grekz
func trap(height []int) int {
    l, r, ml, mr, res := 0, len(height) - 1, 0, 0, 0
    for ( l < r ) {
        if ( height[r] > height[l] ) {
            if ( ml < height[l] ) {
                ml = height[l]
            }else{
                res += ml - height[l]
            }
            l++
        }else{
            if ( mr < height[r] ) {
                mr = height[r]
            }else{
                res += mr - height[r]
            }
            r--
        }
    }
    return res
}
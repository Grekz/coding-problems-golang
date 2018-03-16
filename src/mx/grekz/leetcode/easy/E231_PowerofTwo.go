package easy

func isPowerOfTwo(n int) bool {
    if(n < 1 ) {return false}
    hasOne := false
    for i := 0 ; i < 32; i++{
        if( (n & 1) == 1 ){
            if(hasOne) {
                return false
            } else {
                hasOne = true
            }
        }
        n >>= 1
    }
    return true
}
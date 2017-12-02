package easy

import "math"

func findMedianSortedArrays(n1 []int, n2 []int) float64 {
    a := len(n1)
    b := len(n2)
    if a > b {
        return findMedianSortedArrays(n2, n1)
    }
    c := a + 1 + b
    
    l := 0
    h := a
    for l <= h {
        m1  := ( l + h ) / 2
        m2  := c / 2 - m1
        fmt.Println(m2)
        mla := math.MinInt64
        if m1 != 0 { mla = n1[m1-1] }
        mlb := math.MinInt64
        if m2 != 0 { mlb = n2[m2-1] }
        mra := math.MaxInt64
        if m1 != a { mra = n1[m1] }
        mrb := math.MaxInt64
        if m2 != b { mrb = n2[m2] }
        
        
        if mla > mrb {
            h = m1 - 1
        }else if mlb > mra {
            l = m1 + 1
        }else{
            max := mla
            if mla < mlb {
                max = mlb
            }
            if ( a + b ) % 2 == 0 {
                if mrb > mra {
                    return float64(max + mra) / float64(2)
                }
                return float64(max + mrb) /  float64(2)
            }
            return float64(max)
        }
    }
    return float64(-1)
}

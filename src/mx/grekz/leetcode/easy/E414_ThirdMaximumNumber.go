package easy

// @author grekz
func thirdMax(nums []int) int {
    if len(nums) < 3 {
        _, r := MinMax(nums)
        return r
    }
    m1,_ := MinMax(nums)
    m2, m3 := m1, m1
    for _, x := range nums {
        if ( x > m1 ){ 
            m3 = m2
            m2 = m1
            m1 = x
        }else if ( x > m2 && x < m1 ){ 
            m3 = m2
            m2 = x
        }else if ( x > m3 && x < m2 ){
            m3 = x
        }
    }
    if m1 > m2 && m2 > m3 {
        return m3
    }
    return m1
}
func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}
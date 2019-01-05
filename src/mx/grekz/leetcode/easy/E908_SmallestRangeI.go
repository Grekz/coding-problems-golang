package easy

// @author grekz
func smallestRangeI(A []int, K int) int {
    min, max := MinMax(A)
    res := max - min - 2 * K
    if 0 > res {
        res = 0
    }
    return res
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
package easy

// @author grekz
func canThreePartsEqualSum(A []int) bool {
    sum := 0
    for _, x := range A {
        sum += x
    }
    if sum % 3 != 0 {
        return false
    }
    sum /= 3
    tmp, cnt := 0, 0
    for _, x := range A {
        tmp += x
        if tmp == sum {
            tmp = 0
            cnt++
        }
    }
    return cnt == 3
}
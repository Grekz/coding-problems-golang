package easy

// @author grekz
func minDeletionSize(A []string) int {
    lenA := len(A) - 1
    lenW := len(A[0])
    res := 0
    for i := 0; i < lenW; i++ {
        for j := 0; j < lenA; j++ {
            if A[j][i] > A[j+1][i] { 
                res++
                j = lenA
            }
        }
    }
    return res
}
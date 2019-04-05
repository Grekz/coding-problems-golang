package easy

// @author grekz
func validMountainArray(A []int) bool {
    n, i := len(A), 0
    for ( i + 1 < n && A[i] < A[i + 1] ) {i+=1}  
    if ( i== 0 || i == n - 1 ) {return false}
    for ( i + 1 < n && A[i] > A[i + 1] ) {i+=1}
    return i == n - 1
}
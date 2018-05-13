package easy

// @author grekz
func flipAndInvertImage(A [][]int) [][]int {
    an := len(A)
    for a := 0; a < an; a++ {
        x := A[a]
        j := len(x)-1
        i := 0
        for ( i <= j) {
            tmp := x[i]
            x[i] = x[j] ^ 1
            x[j] = tmp ^ 1
            i += 1
            j -= 1
        }
    }
    return A
}
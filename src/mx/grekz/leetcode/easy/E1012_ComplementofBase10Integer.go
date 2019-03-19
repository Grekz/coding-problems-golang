package easy

// @author grekz
func bitwiseComplement(N int) int {
    a := 1
    for ( a < N ) { a = a * 2 + 1 }
    return a - N
}
package easy

func countPrimes(n int) int {
    if n < 3 { return 0 }
    if n == 3 { return 1 }
    tmp := make( []bool, n+1 )
    for i := 2; i * i <= n; i++ {
        if !tmp[i] {
            for j := i*2; j <= n; j +=i {
                tmp[j] = true
            }
        }
    }
    res := 0
    for i:=2; i < n; i++{
        if !tmp[i] {
            res += 1
        }
    }
    return res
}
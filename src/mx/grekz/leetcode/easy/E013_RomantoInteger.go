package easy

func romanToInt(s string) int {
    m := map[byte]int {
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }
    res := 0
    for i := 0; i < len(s) - 1; i++ {
        cur :=  m[s[i]]
        if cur <  m[s[ i + 1 ]] {
            res -= cur
        }else{
            res += cur
        }
    }
    return res + m[s[ len(s) - 1 ]]
}
package easy

func isIsomorphic(s string, t string) bool {
    n := len(s)
    if n != len(t) {
        return false
    }
    a := make([]int, 256)
    b := make([]int, 256)
    for i := range a {
        a[i] = -1
        b[i] = -1
    }
    for i, si := range s {
        ti := int(t[i])
        if ( a[int(si)] != b[ti] ) { return false }
        a[int(si)] = i 
        b[ti] = i
    }
    return true
}
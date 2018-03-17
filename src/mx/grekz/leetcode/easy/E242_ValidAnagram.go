package easy

// @author grekz
func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {return false}
    tmp := make([]int, 26,26)
    for _, c := range t {
        tmp[int(c - 'a')]+=1
    }
    for _, c := range s {
        if ( tmp[int(c - 'a')] == 0 ) {
            return false
        }
        tmp[int(c - 'a')] -=1
    }
    return true;
}
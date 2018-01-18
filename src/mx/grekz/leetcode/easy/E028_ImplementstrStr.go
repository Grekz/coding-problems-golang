package easy

func strStr(haystack string, needle string) int {
    return str_str_helper(haystack,needle,0)    
}
func str_str_helper(haystack string, needle string, idx int) int {
    nLen := len(needle)
    if ( len(haystack)-idx < nLen ) { return -1 }
    if ( haystack[idx:(idx+nLen)] == needle ) { return idx } 
    return str_str_helper(haystack, needle, idx+1)
}
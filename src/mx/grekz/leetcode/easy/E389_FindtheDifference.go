package easy

// @author grekz
func findTheDifference(s string, t string) byte {
    r := 0
    for _, x := range s {
        r ^= int(x)
    }
    for _, x := range t {
        r ^= int(x)
    }
    return byte(r)
}
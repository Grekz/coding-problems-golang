package easy

// @author grekz
func numUniqueEmails(emails []string) int {
    s := make(map[string]bool)
    for _, e := range emails {
        parts := strings.Split(e, "@")
        a := strings.Replace(strings.Split(parts[0], "+")[0], ".", "", -1)
        b := parts[len(parts) - 1]
        s[a+b] = true
    }
    return len(s)
}
// func numUniqueEmails(emails []string) int {
//     s := make(map[string]bool)
//     for _, e := range emails {
//         parts := regexp.MustCompile("[@+]").Split(e, -1)
//         a := regexp.MustCompile("\\.").ReplaceAllString(parts[0],"")
//         b := parts[len(parts) - 1]
//         s[a+b] = true
//     }
//     return len(s)
// }
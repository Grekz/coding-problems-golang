package easy

// @author grekz
import (
    "strconv"
    "strings"
)
func subdomainVisits(cpdomains []string) []string {
    tmp := make(map[string]int)
    for _, dom := range cpdomains {
        cur := strings.Split(dom, " ")
        count, _ := strconv.Atoi(cur[0])
        sd := cur[1]
        tmp[sd] = getOrZero(tmp, sd) + count
        idx := strings.Index(sd, ".")
        for idx > -1 {
            sd = sd[idx+1:]
            tmp[sd] = getOrZero(tmp, sd) + count
            idx = strings.Index(sd, ".")
        }
    }
    res := make([]string, 0)
    for k, v := range tmp {
        res = append(res, strconv.Itoa(v) + " " + k)
    }
    return res
}
func getOrZero(m map[string]int, k string) int{
    if x, ok := m[k]; ok {
        return x
    }
    return 0
}
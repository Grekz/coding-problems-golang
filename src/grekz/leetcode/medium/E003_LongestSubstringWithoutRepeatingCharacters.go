package easy

func lengthOfLongestSubstring(s string) int {
	res, prev := 0, 0
	dic := make(map[rune]int)
	for i, x := range s {
		if v, ok := dic[x]; ok && v >= prev {
			if d := i - prev; d > res {
				res = d
			}
			prev = v + 1
		}
		dic[x] = i
	}
	if r := len(s) - prev; r > res {
		return r
	}
	return res
}

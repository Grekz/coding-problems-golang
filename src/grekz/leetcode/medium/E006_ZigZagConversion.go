package medium

/**
https://leetcode.com/problems/zigzag-conversion
*/
func convert(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	res := make([]string, numRows)
	i, dir := 0, 0
	for _, x := range s {
		res[i] += string(x)
		if i == 0 {
			dir = 1
		} else if i == numRows-1 {
			dir = -1
		}
		i += dir
	}
	return strings.Join(res, "")
}

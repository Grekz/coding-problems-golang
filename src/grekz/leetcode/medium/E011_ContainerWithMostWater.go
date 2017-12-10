package medium

func maxArea(height []int) int {
    le := len(height)
    if le < 2 { return 0 } 
    l, h, res := 0, le-1, 0
    for l < h {
        min := height[l]
        if height[h] < min {
            min = height[h]
        }
        area := min * ( h - l )
        if res < area {
            res = area
        }
        if height[l] < height[h] {
            l += 1 
        } else {
            h -= 1
        }
    }
    return res
}
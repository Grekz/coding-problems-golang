package easy

// @author grekz
func numRookCaptures(board [][]byte) int {
    rows, cols, ans, posR, posC := len(board), len(board[0]), 0, 0, 0
    _R, _p, _dot := byte('R'), byte('p'), byte('.')
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == _R {
                posR = i
                posC = j
                break
            }
        }
    }
    for i := posR + 1; i < rows; i++ {
        if board[i][posC] != _dot {
            if board[i][posC] == _p {
                ans += 1
            }
            break
        }
    }
    
    for i := posR - 1; i >= 0; i-- {
        if board[i][posC] != _dot {
            if board[i][posC] == _p {
                ans += 1
            }
            break
        }
    }
    for i := posC + 1; i < rows; i++ {
        if board[posR][i] != _dot {
            if board[posR][i] == _p {
                ans += 1
            }
            break
        }
    }
    for i := posC - 1; i >= 0; i-- {
        if board[posR][i] != _dot {
            if board[posR][i] == _p {
                ans += 1
            }
            break
        }
    }
    return ans
    
}
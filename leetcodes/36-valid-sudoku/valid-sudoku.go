func isValidSudoku(board [][]byte) bool {
    seen := [3][9][9]bool{} // Combined storage for rows, cols, and grids

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if cell := board[r][c]; cell != '.' {
                digit := cell - '1' 
                
                if seen[0][r][digit] || seen[1][c][digit] || seen[2][r/3*3+c/3][digit] {
                    return false
                }
                
                seen[0][r][digit] = true
                seen[1][c][digit] = true
                seen[2][r/3*3+c/3][digit] = true
            }
        }
    }
    return true
}

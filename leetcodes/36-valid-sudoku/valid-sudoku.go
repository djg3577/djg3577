func isValidSudoku(board [][]byte) bool {
    rows := [9][9]bool{}
    cols := [9][9]bool{}
    grids := [3][3][9]bool{}

    for r:=0; r<9; r++{
        for c:=0; c<9 ;c++{
            cell := board[r][c]

            if cell == '.'{
                continue
            }

            digit := int(cell) - 49

            if rows[r][digit] || 
               cols[c][digit] || 
               grids[r/3][c/3][digit]{
                return false
            }

            rows[r][digit] = true
            cols[c][digit] = true
            grids[r/3][c/3][digit] = true
        }
    }
    return true
}
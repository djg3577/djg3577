func isValidSudoku(board [][]byte) bool {
    seen := [27]int{} // 9 for rows, 9 for columns, 9 for grids

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if cell := board[r][c]; cell != '.' {
                bit := 1 << (cell - '1') // Create a bitmask for the digit '1'-'9' (0-8)
                grid := (r/3)*3 + c/3    // Unique grid index from 0 to 8

                rowIndex := r
                colIndex := 9 + c
                gridIndex := 18 + grid

                if seen[rowIndex]&bit != 0 || seen[colIndex]&bit != 0 || seen[gridIndex]&bit != 0 {
                    return false
                }

                seen[rowIndex] |= bit
                seen[colIndex] |= bit
                seen[gridIndex] |= bit
            }
        }
    }
    return true
}

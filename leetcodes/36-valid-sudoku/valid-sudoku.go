func isValidSudoku(board [][]byte) bool {
    rowSet := make([]map[byte]struct{}, 9)
    colSet := make([]map[byte]struct{}, 9)
    squareSet := make(map[[2]int]map[byte]struct{}) // Map with tuple-like keys for squares

    // Initialize each set
    for i := 0; i < 9; i++ {
        rowSet[i] = make(map[byte]struct{})
        colSet[i] = make(map[byte]struct{})
    }

    // Initialize square sets for all possible (r/3, c/3) combinations
    for r := 0; r < 3; r++ {
        for c := 0; c < 3; c++ {
            squareSet[[2]int{r, c}] = make(map[byte]struct{})
        }
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            element := board[r][c]
            if element == '.' {
                continue
            }

            // Calculate the square index as a tuple-like key
            squareIndex := [2]int{r / 3, c / 3}

            // Check if the element is already in the corresponding row, column, or square
            if _, found := rowSet[r][element]; found {
                return false
            }
            if _, found := colSet[c][element]; found {
                return false
            }
            if _, found := squareSet[squareIndex][element]; found {
                return false
            }

            // Add the element to the corresponding row, column, and square
            rowSet[r][element] = struct{}{}
            colSet[c][element] = struct{}{}
            squareSet[squareIndex][element] = struct{}{}
        }
    }
    return true
}

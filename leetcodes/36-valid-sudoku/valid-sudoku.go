func isValidSudoku(board [][]byte) bool {
    rowSet := make([]map[byte]struct{}, 9)
    colSet := make([]map[byte]struct{}, 9)
    squareSet := make([]map[byte]struct{}, 9)

    // Initialize each set
    for i := 0; i < 9; i++ {
        rowSet[i] = make(map[byte]struct{})
        colSet[i] = make(map[byte]struct{})
        squareSet[i] = make(map[byte]struct{})
    }

    for r, array := range board {
        for c, element := range array {
            if element == '.' {
                continue
            }
            squareIndex := (r/3)*3 + c/3

            if _, found := rowSet[r][element]; found {
                return false
            }
            if _, found := colSet[c][element]; found {
                return false
            }
            if _, found := squareSet[squareIndex][element]; found {
                return false
            }

            rowSet[r][element] = struct{}{}
            colSet[c][element] = struct{}{}
            squareSet[squareIndex][element] = struct{}{}
        }
    }
    return true
}
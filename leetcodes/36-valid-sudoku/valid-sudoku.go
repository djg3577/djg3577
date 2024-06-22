import "fmt"

func isValidSudoku(board [][]byte) bool {

    rowSet := make([]map[byte]struct{}, 9)
    colSet := make([]map[byte]struct{}, 9)
    squareSet := make([]map[byte]struct{}, 9)

    for i := 0; i < 9; i++{
        rowSet[i] = make(map[byte]struct{})
        colSet[i] = make(map[byte]struct{})
        squareSet[i] = make(map[byte]struct{})
    }

    for r, array := range board{
        for c, element := range array{
            if element == '.' {
                continue
            }

            if _, found := rowSet[r][element]; found{
                return false
            }
            if _, found := colSet[c][element]; found{
                return false
            }
            if _, found := squareSet[(r/3)*3 + c/3][element]; found {
                return false
            }

            rowSet[r][element] = struct{}{}
            colSet[c][element] = struct{}{}
            fmt.Println(r/3)
            fmt.Println((r/3)*3)
            squareSet[(r/3)*3 + c/3][element] = struct{}{}
        }
    }
    return true

}
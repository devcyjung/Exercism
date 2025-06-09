package pascal

func Triangle(n int) [][]int {
	switch {
    case n < 1:
        return nil
    case n == 1:
        return [][]int{{1}}
    default:
        triangle := Triangle(n - 1)
        lastRow := triangle[len(triangle) - 1]
        newRow := make([]int, 0, len(lastRow) + 1)
        newRow = append(newRow, 1)
        for i, v := range lastRow {
            if i == 0 {
                continue
            }
            newRow = append(newRow, lastRow[i - 1] + v)
        }
        newRow = append(newRow, 1)
        return append(triangle, newRow)
    }
}
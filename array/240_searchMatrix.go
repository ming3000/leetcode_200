package array

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	} // if>
	rowCount, columnCount := len(matrix), len(matrix[0])
	row, column := 0, columnCount-1
	for row < rowCount && column >= 0 {
		if target == matrix[row][column] {
			return true
		} else if target < matrix[row][column] {
			column--
		} else {
			row++
		}
	} // for>

	return false
}

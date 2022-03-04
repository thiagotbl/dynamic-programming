package gridtravelertab

func GridTraveler(rows, columns int) int {
	if rows <= 1 || columns <= 1 {
		return 1
	}

	grid := makeGrid(rows+1, columns+1)
	grid[1][1] = 1

	for row := 0; row <= rows; row++ {
		for column := 0; column <= columns; column++ {
			incrementCell(grid, row, column+1, grid[row][column])
			incrementCell(grid, row+1, column, grid[row][column])
		}
	}

	return grid[rows][columns]
}

func makeGrid(rows, columns int) [][]int {
	grid := make([][]int, rows)

	for i := range grid {
		grid[i] = make([]int, columns)
	}

	return grid
}

func incrementCell(grid [][]int, rowIndex, columnIndex, value int) {
	if rowIndex >= len(grid) || columnIndex >= len(grid[rowIndex]) {
		return
	}

	grid[rowIndex][columnIndex] += value
}

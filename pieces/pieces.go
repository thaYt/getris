package pieces

func PieceI() [][]int {
	return [][]int{
		{1, 1, 1, 1},
	}
}

func PieceJ() [][]int {
	return [][]int{
		{2, 2, 2},
		{0, 0, 2},
	}
}

func PieceL() [][]int {
	return [][]int{
		{3, 3, 3},
		{3, 0, 0},
	}
}

func PieceO() [][]int {
	return [][]int{
		{4, 4},
		{4, 4},
	}
}

func PieceS() [][]int {
	return [][]int{
		{0, 5, 5},
		{5, 5, 0},
	}
}

func PieceT() [][]int {
	return [][]int{
		{6, 6, 6},
		{0, 6, 0},
	}
}

func PieceZ() [][]int {
	return [][]int{
		{7, 7, 0},
		{0, 7, 7},
	}
}

func RotateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	newMatrix := make([][]int, cols)
	for i := 0; i < cols; i++ {
		newMatrix[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[j][rows-i-1] = matrix[i][j]
		}
	}

	return newMatrix
}

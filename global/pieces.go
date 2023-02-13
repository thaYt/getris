package global

func PieceI() [][]int {
	return [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}
}

func PieceJ() [][]int {
	return [][]int{
		{0, 0, 0},
		{2, 2, 2},
		{0, 0, 2},
	}
}

func PieceL() [][]int {
	return [][]int{
		{0, 0, 0},
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
		{0, 0, 0},
		{0, 5, 5},
		{5, 5, 0},
	}
}

func PieceT() [][]int {
	return [][]int{
		{0, 0, 0},
		{6, 6, 6},
		{0, 6, 0},
	}
}

func PieceZ() [][]int {
	return [][]int{
		{0, 0, 0},
		{7, 7, 0},
		{0, 7, 7},
	}
}

func RotateMatrix(matrix [][]int) [][]int {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

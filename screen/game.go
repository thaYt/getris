package screen

import (
	"math/rand"
	"thaYt/getris/global"
	"thaYt/getris/pieces"
	"thaYt/getris/states"
	_ "time"
)

var (
	board         [10][20]piece
	dropTime      = 0.75
	Score         int64
	dropped       bool
	droppingPiece bool
	fullDrop      bool
	dropInf       dropPosition
	toDrop        []int
	holdingPiece  bool
	currentPiece  int
	heldPiece     int
	dynScore      int
	level         int32
	rotations     int
)

type piece struct {
	id       int
	dropping bool
}

type dropPosition struct {
	x, y, w, h int
}

func renderGame() {
	for alen, a := range board {
		for blen, b := range a {
			if b.id != 0 {
				drawRectPiece(alen, blen, int32(b.id))
			} else if b.dropping {
				drawRectPieceGrayscale(alen, blen)
			}
		}
	}
}

func renderGameGrayscale() {
	for alen, a := range board {
		for blen, b := range a {
			if b.id != 0 {
				drawRectPieceGrayscale(alen, blen)
			}
		}
	}
}

func checkScore() {
	if dynScore >= 1000 {
		dynScore -= 1000
		upLevel()
	}
}

func upLevel() {
	if dropTime == 0.1 {
		return
	}
	dropTime -= 0.05
	level++
}

func dropPiece() {
	dropped = true
}

func incScore(a int) {
	dynScore += a
	Score += int64(a)
}

func reDrop() {
	dropped = true
	fullDrop = true
}

func checkLineCleared() {
	if !droppingPiece {
		var lines []int
		var aPieces [][]piece
		for j := 0; j < len(board[0]); j++ {
			var line []piece
			for i := 0; i < len(board); i++ {
				if board[i][j].id != 0 {
					line = append(line, board[i][j])
				}
			}
			if len(line) == 10 {
				incScore(100)
				lines = append(lines, j)
			}
		}

		for _, line := range lines {
			for i := range board {
				if board[i][line].id != 0 {
					board[i][line].id = 0
				}
			}
		}

		for j := 0; j < len(board[0]); j++ {
			var line []piece
			for i := 0; i < len(board); i++ {
				line = append(line, board[i][j])
			}

			if amountOfEmptyPieces(line) == 10 {
				continue
			}

			if len(line) > 0 {
				aPieces = append(aPieces, line)
			}
		}

		resetBoard()

		for j := 0; j < len(board[0]); j++ {
			// line := []piece{}
			offset := len(board[0]) - len(aPieces)
			if offset < j {
				continue
			}
			for b, c := range aPieces {
				for pi, pc := range c {
					board[pi][b+offset] = pc
				}
			}
		}
	}
}

func amountOfEmptyPieces(a []piece) (r int) {
	for _, p := range a {
		if p.id == 0 {
			r++
		}
	}
	return r
}

func resetBoard() {
	for b, c := range board {
		for d := range c {
			board[b][d].id = 0
		}
	}
}

func colorBoard() {
	for j := 0; j < len(board[0]); j++ {
		for i := 0; i < len(board); i++ {
			if rand.Int()%10 == 0 {
				board[i][j].id = 0
				continue
			}

			board[i][j].id = rand.Intn(250)
		}
	}
}

func movePieceDown() bool {
	newBoard := board
	for j := 19; j >= 0; j-- {
		for i := 0; i < 10; i++ {
			if board[i][j].dropping {
				if j == 19 || (j < 19 && !board[i][j+1].dropping && board[i][j+1].id != 0) {
					return true
				}
				// move the dropping piece down one row
				newBoard[i][j+1].id = board[i][j].id
				newBoard[i][j+1].dropping = true
				newBoard[i][j].id = 0
				newBoard[i][j].dropping = false
			}
		}
	}
	dropInf.y--
	board = newBoard
	return false
}

func movePieceLeft() bool {
	newBoard := board
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if board[i][j].dropping {
				if i == 0 || (i-1 >= 0 && !board[i-1][j].dropping && board[i-1][j].id != 0) {
					return true
				}

				// move the dropping piece left one column
				newBoard[i-1][j].id = board[i][j].id
				newBoard[i-1][j].dropping = true
				newBoard[i][j].id = 0
				newBoard[i][j].dropping = false
			}
		}
	}
	board = newBoard
	dropInf.x--
	return false
}

func movePieceRight() bool {
	newBoard := board
	for i := 9; i >= 0; i-- {
		for j := 0; j < 20; j++ {
			if board[i][j].dropping {
				if i == 9 || (i+1 < 10 && !board[i+1][j].dropping && board[i+1][j].id != 0) {
					return true
				}

				// move the dropping piece right one column
				newBoard[i+1][j].id = board[i][j].id
				newBoard[i+1][j].dropping = true
				newBoard[i][j].id = 0
				newBoard[i][j].dropping = false
			}
		}
	}
	board = newBoard
	dropInf.x++
	return false
}

func rotatePiece() {
	newBoard := board

	rotated := pieces.RotateMatrix(getPiece(currentPiece))
	for i := 0; i < rotations; i++ {
		rotated = pieces.RotateMatrix(rotated)
	}

	for i, a := range newBoard {
		for j := range a {
			if newBoard[i][j].dropping {
				newBoard[i][j].dropping = false
				newBoard[i][j].id = 0
			}
		}
	}

	dropInf.w, dropInf.h = len(rotated), len(rotated[0])

	for j := 0; j < dropInf.h; j++ {
		for i := 0; i < dropInf.w; i++ {
			if rotated[i][j] != 0 {
				x, y := dropInf.x+i, 19-dropInf.y+j
				if x < 0 || x >= 10 || y < 0 || y >= 20 || newBoard[x][y].id != 0 {
					return
				}
			}
		}
	}

	for j := 0; j < dropInf.h; j++ {
		for i := 0; i < dropInf.w; i++ {
			if rotated[i][j] != 0 {
				x, y := dropInf.x+i, 19-dropInf.y+j
				newBoard[x][y].id = rotated[i][j]
				newBoard[x][y].dropping = true
			}
		}
	}
	board = newBoard
	rotations++
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func startDrop(tet [][]int) {
	a := len(tet)
	if a == 4 {
		a = 2
	} else if a == 2 {
		a = 4
	} else {
		a = 3
	}
	for j := 0; j < len(tet[0]); j++ {
		for i := 0; i < len(tet); i++ {
			if board[i+a][j].id != 0 {
				dropped = true
				global.CurrentMenu = states.DeathScreen
			}
			board[i+a][j].id = tet[i][j]
			board[i+a][j].dropping = true
		}
	}
	for d, b := range board {
		for c := range b {
			if board[d][c].id == 0 && board[d][c].dropping {
				board[d][c].dropping = false
			}
		}
	}
	dropInf.x = a
	dropInf.y = 19
	droppingPiece = true
}

func getPiece(a int) [][]int {
	switch a {
	case 0:
		return pieces.PieceI()
	case 1:
		return pieces.PieceJ()
	case 2:
		return pieces.PieceL()
	case 3:
		return pieces.PieceO()
	case 4:
		return pieces.PieceS()
	case 5:
		return pieces.PieceT()
	case 6:
		return pieces.PieceZ()
	}
	panic("HUH")
}

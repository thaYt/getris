package screen

import (
	"math/rand"
	"thaYt/getris/global"
	_ "time"
)

var (
	Board         [10][20]piece
	SecondDrop    = 1.0
	Score         int64
	dropped       bool
	DroppingPiece bool
	level         int32
	fullDrop      bool
	dropInf       dropPosition
)

type piece struct {
	id       int
	dropping bool
}

type dropPosition struct {
	x int
	y int
	w int
	h int
}

func PlayGame() {
}

func renderGame() {
	for alen, a := range Board {
		for blen, b := range a {
			// c := Board[alen][blen]
			if b.id != 0 {
				drawRectPiece(alen, blen, int32(b.id), b)
			}
		}
	}
}

func renderGameGrayscale() {
	for alen, a := range Board {
		for blen, b := range a {
			if b.id != 0 {
				drawRectPieceGrayscale(alen, blen)
			}
		}
	}
}

func upLevel() {
	if SecondDrop == 0.5 {
		return
	}
	SecondDrop -= 0.1
	level++
}

func dropPiece() {
	dropped = true
}

func reDrop() {
	dropped = true
	fullDrop = true
}

func checkLineCleared() {
	if !DroppingPiece {
		var lines []int
		var aPieces [][]piece
		for j := 0; j < len(Board[0]); j++ {
			line := []piece{}
			for i := 0; i < len(Board); i++ {
				if Board[i][j].id != 0 {
					line = append(line, Board[i][j])
				}
			}
			if len(line) == 10 {
				lines = append(lines, j)
			}
		}

		Score += 100 * int64(len(aPieces))

		for _, line := range lines {
			for i := range Board {
				if Board[i][line].id != 0 {
					Board[i][line].id = 0
				}
			}
		}

		for j := 0; j < len(Board[0]); j++ {
			line := []piece{}
			for i := 0; i < len(Board); i++ {
				line = append(line, Board[i][j])
			}

			if amountOfEmptyPieces(line) == 10 {
				continue
			}

			if len(line) > 0 {
				aPieces = append(aPieces, line)
			}
		}

		resetBoard()

		for j := 0; j < len(Board[0]); j++ {
			// line := []piece{}
			offset := len(Board[0]) - len(aPieces)
			if offset < j {
				continue
			}
			for b, c := range aPieces {
				for pi, pc := range c {
					Board[pi][b+offset] = pc
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
	for b, c := range Board {
		for d := range c {
			Board[b][d].id = 0
		}
	}
}

func colorBoard() {
	for j := 0; j < len(Board[0]); j++ {
		for i := 0; i < len(Board); i++ {
			if rand.Int()%10 == 0 {
				Board[i][j].id = 0
				continue
			}

			Board[i][j].id = rand.Intn(250)
		}
	}
}

func movePieceDown() bool {
	newBoard := Board
	for j := 19; j >= 0; j-- {
		for i := 0; i < 10; i++ {
			if Board[i][j].dropping {
				if j+1 == 20 && Board[i][j].id != 0 {
					return true
				}
				if !Board[i][j+1].dropping && Board[i][j+1].id != 0 {
					return true
				}
				// move the dropping piece down one row
				newBoard[i][j+1].id = Board[i][j].id
				newBoard[i][j+1].dropping = true
				newBoard[i][j].id = 0
				newBoard[i][j].dropping = false
			}
		}
	}
	Board = newBoard
	return false
}

func movePieceLeft() bool {
	newBoard := Board
	for j := 10; j < len(Board[0]); j++ {
		for i := 0; i < len(Board); i++ {
			if Board[i][j].dropping {
				if !Board[i-1][j].dropping && Board[i-1][j].id != 0 {
					return true
				}
				newBoard[i][j] = Board[i][j]
			}
		}
	}
	Board = newBoard
	return false
}

func movePieceRight() bool {
	var newBoard = Board
	for j := 10; j < len(Board[0]); j++ {
		for i := 0; i < len(Board); i++ {
			if Board[i][j].dropping {
				if !Board[i-1][j].dropping && Board[i-1][j].id != 0 {
					return true
				}
				newBoard[i][j] = Board[i][j]
			}
		}
	}
	Board = newBoard
	return false
}

func rotatePiece() {
	// global.RotateMatrix(Board[dropInf.x : dropInf.w+dropInf.x][dropInf.y : dropInf.h+dropInf.y])
}

func startDrop(tet [][]int) {
	a := len(tet)
	if a == 3 {
		a -= 1
	}
	a = a/2 + 2
	for j := 0; j < len(tet[0]); j++ {
		for i := 0; i < len(tet); i++ {
			if Board[i+a][j].id != 0 {
				dropped = true
				global.CurrentMenu = global.DeathScreen
			}
			Board[i+a][j].id = tet[i][j]
			Board[i+a][j].dropping = true
		}
	}
	DroppingPiece = true
}

func getRandomPiece() [][]int {
	switch rand.Intn(7) {
	case 0:
		return global.PieceI()
	case 1:
		return global.PieceJ()
	case 2:
		return global.PieceL()
	case 3:
		return global.PieceO()
	case 4:
		return global.PieceS()
	case 5:
		return global.PieceT()
	case 6:
		return global.PieceZ()
	}
	panic("HUH")
}

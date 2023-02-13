package screen

import (
	"fmt"
	"thaYt/getris/global"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleKey(number int) {
	switch number { // global
	case 9:
		if global.Vsync == 0 {
			global.Vsync = 1
			fmt.Println("vsync enabled")
		} else {
			global.Vsync = 0
			fmt.Println("vsync disabled")
		}
		sdl.GLSetSwapInterval(global.Vsync)
	}
	switch global.CurrentMenu {
	case global.TitleScreen:
		switch number {
		case 41:
			global.Running = false
		case 82: // menu up
			if titleOpt == 0 {
				titleOpt++
			} else {
				titleOpt--
			}
		case 81: // menu down
			titleOpt++
		case 44, 40: // click
			selected := titleOpt % 2
			switch selected {
			case 0:
				global.CurrentMenu = global.GameScreen
			case 1:
				global.CurrentMenu = global.InfoScreen
			}
		}
	case global.GameScreen:
		switch number {
		case 82: // up | rotate
			rotatePiece()
		case 80: // left | move piece left
			movePieceLeft()
		case 81: // down | push piece down
			dropPiece()
		case 79: // right | move piece right
			movePieceRight()
		case 44: // space | hard drop
			reDrop()
		case 225, 229: // l/r shift | hold piece
		case 41:
			global.CurrentMenu = global.DeathScreen
		}
	case global.DeathScreen:
		switch number {
		case 44, 40, 41:
			global.CurrentMenu = global.TitleScreen
		}
	case global.InfoScreen:
		global.CurrentMenu = global.TitleScreen
	}
}

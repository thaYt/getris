package screen

import (
	"fmt"
	"math/rand"
	"strconv"
	"thaYt/getris/fonts"
	"thaYt/getris/global"
	"thaYt/getris/states"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	titleOpt int
	started  bool
	lastDrop time.Time
	needHold bool
	alrHeld  bool
)

func DrawScreen(number int) {
	fillRectColor(&sdl.Rect{X: 0, Y: 0, W: global.WIDTH, H: global.HEIGHT}, global.Black)
	switch number {
	case states.DeathScreen:
		drawDeathScreen()
	case states.TitleScreen:
		drawTitleScreen()
	case states.GameScreen:
		drawGameScreen()
	case states.InfoScreen:
		drawInfoScreen()
	}
	drawFPS()
	global.Renderer.Present()
}

func drawInfoScreen() {
	drawBack()
	drawFont(global.WIDTH/2, 65, fonts.TitleFont, "keys/info", &global.White, true)

	drawFont(175, 125, fonts.SubtitleFont, "global bindings:", &global.White, false)
	drawFont(175, 165, fonts.SubtitleFont, "f to enable vsync", &global.White, false)
	drawFont(175, 195, fonts.SubtitleFont, "esc to leave menu", &global.White, false)

	drawFont(175, 255, fonts.SubtitleFont, "game bindings:", &global.White, false)
	drawFont(175, 295, fonts.SubtitleFont, "left/right to move piece", &global.White, false)
	drawFont(175, 325, fonts.SubtitleFont, "up to rotate piece", &global.White, false)
	drawFont(175, 355, fonts.SubtitleFont, "down to move piece down", &global.White, false)
	drawFont(175, 385, fonts.SubtitleFont, "space to harddrop", &global.White, false)
	drawFont(175, 415, fonts.SubtitleFont, "l/r shift to hold piece", &global.White, false)

	fillRectColor(&sdl.Rect{X: global.XOFF, Y: global.YOFF + global.YRES/2, W: global.XRES, H: 2}, global.White)

	drawFont(175, 550, fonts.SubtitleFont, "Getris was made by me (core) with", &global.White, false)
	drawFont(175, 580, fonts.SubtitleFont, "inspiration from Lightcaster5's NES Tetris", &global.White, false)
	drawFont(175, 610, fonts.SubtitleFont, "remake. I wanted to expand on my Golang", &global.White, false)
	drawFont(175, 640, fonts.SubtitleFont, "skills and make a simple game with the", &global.White, false)
	drawFont(175, 670, fonts.SubtitleFont, "SDL2 library, and I think I did pretty well.", &global.White, false)

	drawFont(175, 710, fonts.SubtitleFont, "Credits:", &global.White, false)
	drawFont(175, 750, fonts.SubtitleFont, "core: development", &global.White, false)
	drawFont(175, 780, fonts.SubtitleFont, "Lightcaster5: inspiration, colors, extra help", &global.White, false)

	drawFooter("press any key to leave")
}

func drawDeathScreen() {
	if started {
		started = false
	}
	drawBack()

	renderGameGrayscale()

	drawFont(global.WIDTH/2, 65, fonts.ButtonFont, "you died!", &global.White, true)
	drawFont(global.WIDTH/2, 100, fonts.ButtonFont, fmt.Sprintf("score: %d", Score), &global.White, true)
}

func drawTitleScreen() {
	drawBack()
	drawFont(global.WIDTH/2, 65, fonts.TitleFont, "getris", &global.White, true)
	drawFont(global.WIDTH/2, 115, fonts.SubtitleFont, "made by core!", &global.White, true)

	// this sucks because it relies on there being 2 options only. please change this eventually
	switch titleOpt % 2 {
	case 0:
		drawFont(global.WIDTH/2, 600, fonts.ButtonFont, "[ start ]", &global.White, true)
		drawFont(global.WIDTH/2, 675, fonts.ButtonFont, "controls/info", &global.Gray, true)
	case 1:
		drawFont(global.WIDTH/2, 600, fonts.ButtonFont, "start", &global.Gray, true)
		drawFont(global.WIDTH/2, 675, fonts.ButtonFont, "[ controls/info ]", &global.White, true)
	}

	drawFooter("up/down to select, enter/space to choose")
}

func drawGameScreen() {
	// init game
	if !started {
		board = [10][20]piece{}
		lastDrop = time.Now()
		started = true
		toDrop = []int{}
		heldPiece = 0
		currentPiece = 0
		for a := 0; a < 6; a++ {
			toDrop = append(toDrop, rand.Intn(7))
		}
	}
	// piece dropping
	if !droppingPiece {
		if !needHold {
			currentPiece = toDrop[0]
			toDrop = toDrop[1:]
			toDrop = append(toDrop, rand.Intn(7))
		}
		pce := getPiece(currentPiece)
		rotations = 0
		startDrop(pce)
		dropInf.w, dropInf.h = len(pce[0]), len(pce)

		checkLineCleared()
		if needHold {
			needHold = false
		}
	}
	if time.Since(lastDrop).Milliseconds() >= int64(dropTime*1000) || dropped {
		lastDrop = time.Now()
		dropped = false
		if fullDrop {
			fullDrop = false
			for !movePieceDown() {
			}
			killDrop()
		}
		if movePieceDown() {
			killDrop()
		}
	}
	checkLineCleared()
	checkScore()

	drawBack()
	renderGame()

	// score
	drawFont(global.XOFF/2, 50, fonts.SubtitleFont, "score", &global.White, true)
	drawFont(global.XOFF/2, 75, fonts.SubtitleFont, strconv.FormatInt(Score, 10), &global.White, true)

	// held
	drawFont(global.XOFF/2, 150, fonts.SubtitleFont, "held piece", &global.White, true)
	if holdingPiece {
		renderHoldPiece()
	}

	r := sdl.Rect{
		X: global.XOFF + global.XRES + 2,
		Y: global.YOFF,
		W: global.XOFF,
		H: global.HEIGHT/7 + global.YOFF,
	}
	b := getBorder(r, 2)
	fillRectColor(&b, global.White)
	fillRectColor(&r, sdl.Color{R: 25, G: 25, B: 25, A: 255})

	// next pieces
	renderNextPieces()
}

func holdPiece() {
	alrHeld = true
	droppingPiece = true
	if holdingPiece {
		currentPiece, heldPiece = heldPiece, currentPiece
		needHold = true
		alrHeld = true
	} else {
		holdingPiece = true
		heldPiece = currentPiece
	}
	droppingPiece = false
	for i, a := range board {
		for j, b := range a {
			if b.id != 0 && b.dropping {
				board[i][j].id, board[i][j].dropping = 0, false
			}
		}
	}
}

func renderHoldPiece() {
	// don't you love int conversions?
	pce := getPiece(int(heldPiece))
	for p, a := range pce {
		for j, b := range a {
			if b == 0 {
				continue
			}
			color := global.Ca[int32(heldPiece)%7+level][0]
			fillRectColor(
				&sdl.Rect{
					X: global.XOFF/4 + (global.XRES / 15 * int32(p)) + 1,
					Y: int32(global.HEIGHT/len(toDrop)) + 25 + int32(global.XRES/15*j) + 1,
					W: global.XRES/15 - 2,
					H: global.YRES/30 - 2,
				},
				color,
			)
			color = global.Ca[int32(heldPiece)%7+level][1]
			fillRectColor(
				&sdl.Rect{
					X: global.XOFF/4 + (global.XRES / 15 * int32(p)) + 3,
					Y: int32(global.HEIGHT/len(toDrop)) + 25 + int32(global.XRES/15*j) + 3,
					W: global.XRES/15 - 6,
					H: global.YRES/30 - 6,
				},
				color,
			)
		}
	}
}

func renderNextPieces() {
	// don't you love int conversions?
	for i := 0; i < len(toDrop); i++ {
		i := int32(i)
		pce := getPiece(int(toDrop[i]))
		for p, a := range pce {
			for j, b := range a {
				if b == 0 {
					continue
				}
				color := global.Ca[int32(toDrop[i])%7+level][0]
				fillRectColor(
					&sdl.Rect{
						X: global.XRES + global.XOFF + 50 + (global.XRES / 15 * int32(p)) + 1,
						Y: int32(global.HEIGHT/len(toDrop))*i + 25 + int32(global.XRES/15*j) + 1,
						W: global.XRES/15 - 2,
						H: global.YRES/30 - 2,
					},
					color,
				)
				color = global.Ca[int32(toDrop[i])%7+level][1]
				fillRectColor(
					&sdl.Rect{
						X: global.XRES + global.XOFF + 50 + (global.XRES / 15 * int32(p)) + 3,
						Y: int32(global.HEIGHT/len(toDrop))*i + 25 + int32(global.XRES/15*j) + 3,
						W: global.XRES/15 - 6,
						H: global.YRES/30 - 6,
					},
					color,
				)
			}
		}
	}
}

func killDrop() {
	droppingPiece = false
	for i, By := range board {
		for j := range By {
			board[i][j].dropping = false
		}
	}
}

func drawFPS() {
	drawFont(5, 5, fonts.FPSFont, fmt.Sprintf("FPS: %.2f", global.FPS), &global.White, false)
}

func drawFont(x, y int32, font *ttf.Font, text string, color *sdl.Color, centered bool) {
	surface, err := font.RenderUTF8Blended(text, *color)
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	rect := sdl.Rect{X: x, Y: y, W: surface.W, H: surface.H}
	if centered {
		rect.X = rect.X - surface.W/2
		rect.Y = rect.Y - surface.H/2
	}

	texture, err := global.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	global.Renderer.Copy(texture, nil, &rect)
}

func getBorder(rect sdl.Rect, size int32) sdl.Rect {
	return sdl.Rect{
		X: rect.X - (1 * size),
		Y: rect.Y - (1 * size),
		W: rect.W + (2 * size),
		H: rect.H + (2 * size),
	}
}

func drawFooter(str string) {
	drawFont(global.WIDTH/2, 970, fonts.SubtitleFont, str, &global.White, true)
}

func fillRectColor(rect *sdl.Rect, color sdl.Color) {
	global.Renderer.SetDrawColor(extendColor(color))
	global.Renderer.FillRect(rect)
}

func drawBack() {
	inside := sdl.Rect{X: global.XOFF, Y: global.YOFF, W: global.XRES, H: global.YRES}
	border := getBorder(inside, 2)
	fillRectColor(&border, global.White)
	fillRectColor(&inside, global.Black)
}

func drawRectPiece(boardX, boardY int, id int32) {
	id++
	color := global.Ca[id%7+level][0]
	fillRectColor(
		&sdl.Rect{
			X: getPieceXOffset(0) + (global.XRES / 10 * int32(boardX)) + 1,
			Y: getPieceYOffset(0) + (global.YRES / 20 * int32(boardY)) + 1,
			W: global.XRES/10 - 2,
			H: global.YRES/20 - 2,
		},
		color,
	)

	color = global.Ca[id%7+level][1]
	fillRectColor(
		&sdl.Rect{
			X: (getPieceXOffset(0) + 5) + (global.XRES / 10 * int32(boardX)),
			Y: (getPieceYOffset(0) + 5) + (global.YRES / 20 * int32(boardY)),
			W: global.XRES/10 - 10,
			H: global.YRES/20 - 10,
		},
		color,
	)
}

func drawRectPieceGrayscale(boardX, boardY int) {
	color := sdl.Color{R: 125, G: 125, B: 125, A: 255}
	fillRectColor(
		&sdl.Rect{
			X: getPieceXOffset(0) + (global.XRES / 10 * int32(boardX)) + 1,
			Y: getPieceYOffset(0) + (global.YRES / 20 * int32(boardY)) + 1,
			W: global.XRES/10 - 2,
			H: global.YRES/20 - 2,
		},
		color,
	)
}

func getPieceXOffset(xset int) int32 {
	return int32(global.XOFF + xset*(global.XRES/10))
}

func getPieceYOffset(yset int) int32 {
	return int32(global.YOFF + yset*(global.YRES/20))
}

func extendColor(color sdl.Color) (uint8, uint8, uint8, uint8) {
	return color.R, color.G, color.B, color.A
}

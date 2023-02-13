package screen

import (
	"fmt"
	"strconv"
	"thaYt/getris/fonts"
	"thaYt/getris/global"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	titleOpt int
	started  bool
	lastDrop time.Time
)

func DrawScreen(number int) {
	fillRectColor(&sdl.Rect{X: 0, Y: 0, W: global.WIDTH, H: global.HEIGHT}, global.Black)
	switch number {
	case global.DeathScreen:
		drawDeathScreen()
	case global.TitleScreen:
		drawTitleScreen()
	case global.GameScreen:
		drawGameScreen()
	case global.InfoScreen:
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
	if !started {
		lastDrop = time.Now()
		PlayGame()
		started = true
	}
	if !DroppingPiece {
		startDrop(getRandomPiece())
	}
	if time.Since(lastDrop).Milliseconds() >= int64(SecondDrop*1000) || dropped {
		print("drop!!!!")
		dropped = false
		if movePieceDown() {
			killDrop()
		}
		lastDrop = time.Now()
	}
	// checkLineCleared()

	drawBack()
	renderGame()

	// score
	drawFont(global.XOFF/2, 50, fonts.SubtitleFont, "score", &global.White, true)
	drawFont(global.XOFF/2, 75, fonts.SubtitleFont, strconv.FormatInt(Score, 10), &global.White, true)

	// held
	drawFont(global.XOFF/2, 150, fonts.SubtitleFont, "held piece", &global.White, true)

	// drawBorder(global.XOFF, global.YOFF, global.XRES, global.YRES, 1, global.White)
}

func killDrop() {
	DroppingPiece = false
	for i, By := range Board {
		for j := range By {
			Board[i][j].dropping = false
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

/* doesn't work fsr
func drawBorder(x, y, width, height, borderSize int32, borderColor, insideColor *sdl.Color) {
	border := sdl.Rect{X: x - (1 * borderSize), Y: y - (1 * borderSize), W: width + (2 * borderSize), H: width + (2 * borderSize)}
	inside := sdl.Rect{X: x, Y: y, W: width, H: height}
	global.Renderer.SetDrawColor(borderColor.R, borderColor.G, borderColor.B, borderColor.A)
	global.Renderer.FillRect(&border)
	global.Renderer.SetDrawColor(insideColor.R, insideColor.G, insideColor.B, insideColor.A)
	global.Renderer.FillRect(&inside)
}
*/

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

func drawRectPiece(boardX, boardY int, id int32, a piece) {
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
	if a.dropping {
		fillRectColor(
			&sdl.Rect{
				X: (getPieceXOffset(0) + 5) + (global.XRES / 10 * int32(boardX)),
				Y: (getPieceYOffset(0) + 5) + (global.YRES / 20 * int32(boardY)),
				W: global.XRES/10 - 10,
				H: global.YRES/20 - 10,
			},
			sdl.Color{R: 170, G: 0, B: 0, A: 0},
		)
	}
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

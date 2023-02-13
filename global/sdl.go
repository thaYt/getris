package global

import "github.com/veandco/go-sdl2/sdl"

var (
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Surface  *sdl.Surface

	FPS         float64
	Vsync       int
	CurrentMenu = TitleScreen
	Running bool
)

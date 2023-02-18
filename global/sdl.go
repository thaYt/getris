package global

import (
	"github.com/veandco/go-sdl2/sdl"
	"thaYt/getris/states"
)

var (
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Surface  *sdl.Surface

	FPS         float64
	Vsync       int
	CurrentMenu = states.TitleScreen
	Running     bool
)

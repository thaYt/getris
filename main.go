package main

import (
	"fmt"
	"thaYt/getris/fonts"
	"thaYt/getris/global"
	"thaYt/getris/screen"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	bt := time.Now().UnixMilli()
	fmt.Println("getris loading...")
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	defer ttf.Quit()
	fonts.LoadFonts()
	fmt.Println("loaded fonts...")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	fmt.Println("sdl initiated...")

	var err error
	global.Window, err = sdl.CreateWindow("getris", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		global.WIDTH, global.HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer global.Window.Destroy()
	fmt.Println("sdl window created...")

	global.Renderer, err = sdl.CreateRenderer(global.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer global.Renderer.Destroy()
	fmt.Println("sdl renderer created...")

	fmt.Printf("done! loaded in %dms\n", time.Now().UnixMilli()-bt)

	fmt.Printf("resolution: x: %d, y: %d\n", global.XRES, global.YRES)

	var frameCount int
	startTime := time.Now()

	global.Running = true
	for global.Running {
		frameCount++
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.KeyboardEvent:
				if event.State == 1 {
					keyCode := event.Keysym.Scancode
					fmt.Printf("key pressed: %d\n", keyCode)
					screen.HandleKey(int(keyCode))
				}
			case *sdl.QuitEvent:
				fmt.Printf("Quit: lasted %ss\n", fmt.Sprint(event.Timestamp/1000))
				global.Running = false
			}
		}
		screen.DrawScreen(global.CurrentMenu)
		timeElapsed := time.Since(startTime).Seconds()
		if timeElapsed >= 1 {
			startTime = time.Now()
			global.FPS = float64(frameCount) / timeElapsed
			frameCount = 0
		}
	}
}

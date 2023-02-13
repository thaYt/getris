package fonts

import (
	_ "embed"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//go:embed Ubuntu-Bold.ttf
var ubuntuMediumFile []byte

//go:embed Ubuntu-Light.ttf
var ubuntuLightFile []byte

//go:embed Ubuntu-Medium.ttf
var ubuntuBoldFont []byte

// to use these, run LoadFonts() first  
var (
	SubtitleFont *ttf.Font
	TitleFont    *ttf.Font
	FPSFont      *ttf.Font
	ButtonFont   *ttf.Font
)

func LoadFonts() {
	ff, _ := sdl.RWFromMem(ubuntuLightFile)
	FPSFont, _ = ttf.OpenFontRW(ff, 0, 18)

	bf, _ := sdl.RWFromMem(ubuntuLightFile)
	ButtonFont, _ = ttf.OpenFontRW(bf, 0, 45)

	ulf, _ := sdl.RWFromMem(ubuntuLightFile)
	SubtitleFont, _ = ttf.OpenFontRW(ulf, 0, 24)

	tf, _ := sdl.RWFromMem(ubuntuBoldFont)
	TitleFont, _ = ttf.OpenFontRW(tf, 0, 92)
}

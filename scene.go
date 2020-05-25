package terminus

import(

	"github.com/gdamore/tcell"

)

type Scene struct {

	screen tcell.Screen
	foreground tcell.Color
	background tcell.Color

	entities []Entity

}

func NewScene() *Scene {

	scene := &Scene{
		tcell.NewScreen(),
		WHITE,
		BLACK,
		[]Entity{}
	}

	return scene

}

func NewScene( fg, bg tcell.Color ) *Scene {

	scene := &Scene{
		tcell.NewScreen(),
		fg,
		bg,
		[]Entity{}
	}

	return scene

}
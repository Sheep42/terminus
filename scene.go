package terminus

import(

	"github.com/gdamore/tcell"
	"os"

)

type Scene struct {

	screen tcell.Screen
	foreground tcell.Color
	background tcell.Color

	entities []*Entity

}

func NewScene() (*Scene, error) {

	screen, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}

	scene := &Scene{
		screen,
		BLACK,
		WHITE,
		[]*Entity{},
	}

	return scene, nil

}

func NewSceneCustom( fg, bg tcell.Color ) (*Scene, error) {

	screen, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}

	scene := &Scene{
		screen,
		fg,
		bg,
		[]*Entity{},
	}

	return scene, nil

}

func (scene *Scene) Add( entity *Entity ) {

	scene.entities = append( scene.entities, entity )

}

func (scene *Scene) Init() {

	screen := scene.screen

	screen_style := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screen_style)

	screen.Init()

}

func (scene *Scene) Draw() {

	screen := scene.screen

	screen.Clear()
	screen.Show()

	ev := screen.PollEvent()

	switch ev := ev.(type) {

		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
		default:

	}

}
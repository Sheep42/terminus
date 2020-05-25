package terminus

type Game struct {

	current_scene *Scene

}

func NewGame( current_scene *Scene ) *Game {

	game := &Game{
		current_scene,
	}

	return game

}

func (game *Game) Run() {

	game.current_scene.Init()

// game_loop:
	for {

		game.current_scene.Draw()

	}

}

func (game *Game) ChangeScene( scene *Scene ) {

	game.current_scene = scene

}
package terminus

type Game struct {

	current_scene Scene

}

func NewGame( current_scene Scene ) *Game {

	game := &Game{
		current_scene
	}

	return game

}

func (*Game) Run() {


// game_loop:
	for {



	}

}
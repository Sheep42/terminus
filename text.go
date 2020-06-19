package terminus

type Text struct {

	*Entity
	text string

}

func NewText( x, y int, text string ) *Text {

	t := &Text{
		Entity: NewEntity(x, y),
		text: text,
	}

	return t

}

func (text *Text) Init() {
	text.Entity.Init() // super
}

func (text *Text) Update( delta float64 ) { 
	text.Entity.Update( delta ) // super
}

func (text *Text) Draw() {

	// override
	screen := text.Entity.game.screen
	style := text.Entity.scene.style

	for index, char := range text.text {
		screen.SetContent(text.x + index, text.y, rune(char), nil, style)
	}
	
}

func (text *Text) SetText( new_text string ) {
	text.text = new_text
}

func (text *Text) GetText() string {
	return text.text
}
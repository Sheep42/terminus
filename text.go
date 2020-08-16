package terminus

// Text is a type of Entity which is used
// to render text to the game screen
type Text struct {
	*Entity
	text string
}

// NewText takes an x position, y position, and text
// value and creates a new Text Entity on the screen
func NewText(x, y int, text string) *Text {

	t := &Text{
		Entity: NewEntity(x, y),
		text:   text,
	}

	return t

}

// Init fires during game.Init and can be overridden
func (text *Text) Init() {
	text.Entity.Init() // super
}

// Update fires after the scene update on each pass
// through the game loop, and can be overridden
func (text *Text) Update(delta float64) {
	text.Entity.Update(delta) // super
}

// Draw fires during scene.Draw and can be overridden
func (text *Text) Draw() {

	// override Entity.Draw
	screen := text.Entity.game.screen
	style := text.Entity.scene.style

	for index, char := range text.text {
		screen.SetContent(text.x+index, text.y, rune(char), nil, style)
	}

}

// SetText sets the text value of the Text Entity
func (text *Text) SetText(newText string) {
	text.text = newText
}

// GetText gets the text value of the Text Entity
func (text *Text) GetText() string {
	return text.text
}

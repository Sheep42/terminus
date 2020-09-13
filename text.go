package terminus

// IText is the interface through which custom
// implementations of Text can be created
type IText interface {
	ToEntities(text string) []IEntity
}

// Text is a type of EntityGroup which is used
// to render text to the game screen
type Text struct {
	*EntityGroup
	text string
}

// NewText takes an x position, y position, and text
// value and creates a new Text Entity on the screen
func NewText(x, y int, text string) *Text {

	entities := ToEntities(text)

	t := &Text{
		EntityGroup: NewEntityGroup(x, y, len(text), 1, entities),
		text:        text,
	}

	return t

}

// Update fires after the scene update on each pass
// through the game loop, and can be overridden
func (t *Text) Update(delta float64) {
	t.EntityGroup.Update(delta) // super
}

// ToEntities returns a slice of entities
// representing a given string of text
func ToEntities(text string) []IEntity {

	entities := []IEntity{}

	for index, char := range text {
		// 0,0 starts from top left of the
		// EntityGroup
		entities = append(entities, NewSpriteEntity(index, 0, rune(char)))
	}

	return entities
}

// SetText sets the text value of the Text Entity
func (t *Text) SetText(newText string) {

	t.text = newText
	t.entities = ToEntities(newText)
	t.EntityGroup.SetWidth(len(newText))

}

// GetText gets the text value of the Text Entity
func (t *Text) GetText() string {
	return t.text
}

// GetEntityGroup gets the EntityGroup that contain
// the Text Entities
func (t *Text) GetEntityGroup() *EntityGroup {
	return t.EntityGroup
}

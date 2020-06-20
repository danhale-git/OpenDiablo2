package d2input

// MouseButton represents a traditional 3-button mouse
type MouseButton int

const (
	// MouseButtonLeft is the left mouse button
	MouseButtonLeft MouseButton = iota
	// MouseButtonMiddle is the middle mouse button
	MouseButtonMiddle
	// MouseButtonRight is the right mouse button
	MouseButtonRight

	mouseButtonMin = MouseButtonLeft
	mouseButtonMax = MouseButtonRight
)

// MouseButtonMod represents a "modified" mouse button action. This could mean, for example, ctrl-mouse_left
type MouseButtonMod int

const (
	// MouseButtonLeft is a modified left mouse button
	MouseButtonModLeft MouseButtonMod = 1 << iota
	// MouseButtonModMiddle is a modified middle mouse button
	MouseButtonModMiddle
	// MouseButtonModRight is a modified right mouse button
	MouseButtonModRight
)

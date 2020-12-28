package abfactory

// GUIFactory interface
type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

// WinFactory class
type WinFactory struct{}

// NewWinFactory constructor
func NewWinFactory() *WinFactory {
	return &WinFactory{}
}

// CreateButton method
func (w *WinFactory) CreateButton() Button {
	return NewWinButton()
}

// CreateCheckBox method
func (w *WinFactory) CreateCheckBox() CheckBox {
	return NewWinCheckBox()
}

// MacFactory class
type MacFactory struct {
	_Button   Button
	_CheckBox CheckBox
}

// NewMacFactory constructor
func NewMacFactory() *MacFactory {
	return &MacFactory{}
}

// CreateButton method
func (m *MacFactory) CreateButton() Button {
	return NewMacButton()
}

// CreateCheckBox method
func (m *MacFactory) CreateCheckBox() CheckBox {
	return NewMacCheckBox()
}

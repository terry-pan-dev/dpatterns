package abfactory

import "fmt"

// Button interface
type Button interface {
	Paint()
}

// CheckBox interface
type CheckBox interface {
	Paint()
}

// WinButton class
type WinButton struct{}

// NewWinButton constructor
func NewWinButton() *WinButton {
	return &WinButton{}
}

// Paint method
func (w *WinButton) Paint() {
	fmt.Println("Paint Windows button")
}

// MacButton class
type MacButton struct{}

// NewMacButton constructor
func NewMacButton() *MacButton {
	return &MacButton{}
}

// Paint method
func (m *MacButton) Paint() {
	fmt.Println("Paint Mac button")
}

// WinCheckBox class
type WinCheckBox struct{}

// NewWinCheckBox constructor
func NewWinCheckBox() *WinCheckBox {
	return &WinCheckBox{}
}

// Paint method
func (w *WinCheckBox) Paint() {
	fmt.Println("Paint Windows checkbox")
}

// MacCheckBox class
type MacCheckBox struct{}

// NewMacCheckBox constructor
func NewMacCheckBox() *MacCheckBox {
	return &MacCheckBox{}
}

// Paint method
func (m *MacCheckBox) Paint() {
	fmt.Println("Paint Mac checkbox")
}

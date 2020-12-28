package abfactory

// MyAbstractFactoryAppDemo method
func MyAbstractFactoryAppDemo() {
	var winGui GUIFactory = NewWinFactory()
	winGui.CreateButton().Paint()
	winGui.CreateCheckBox().Paint()
	var macGui GUIFactory = NewMacFactory()
	macGui.CreateButton().Paint()
	macGui.CreateCheckBox().Paint()
}

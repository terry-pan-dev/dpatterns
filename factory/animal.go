package factory

import "fmt"

// Animal interface
// any conconcret animal class must confirm this signature
type Animal interface {
	Speak()
	PreferedAction()
}

// Dog class
// the dog class
type Dog struct{}

// Speak method
// the speak method of dog
func (d *Dog) Speak() {
	fmt.Println("Bak, bak")
}

// PreferedAction method
func (d *Dog) PreferedAction() {
	fmt.Println("Dog prefers running")
}

// Cat class
// the cat class
type Cat struct{}

// Speak method
// the speak method of dog
func (c *Cat) Speak() {
	fmt.Println("Niam, Niam")
}

// PreferedAction method
func (c *Cat) PreferedAction() {
	fmt.Println("Dog prefers sleeping")
}

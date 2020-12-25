package observer

import "fmt"

// Observer interface
type Observer interface {
	Update(int)
}

// Observer1 concrete class one
type Observer1 struct {
	Name string
}

// Observer2 concrete class one
type Observer2 struct {
	Name string
}

// NewObserver1 constructor
func NewObserver1(name string) *Observer1 {
	return &Observer1{Name: name}
}

// Update method
func (o *Observer1) Update(v int) {
	fmt.Printf("Value has been updated %d for observer %s\n", v, o.Name)
}

// NewObserver2 constructor
func NewObserver2(name string) *Observer2 {
	return &Observer2{Name: name}
}

// Update method
func (o *Observer2) Update(v int) {
	fmt.Printf("Value has been updated %d for observer %s\n", v, o.Name)
}

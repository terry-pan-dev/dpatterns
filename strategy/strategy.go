package strategy

import "fmt"

// Strategy interface
type Strategy interface {
	BuildRoute(float32, float32)
}

// Road strategy
type Road struct{}

// Walk strategy
type Walk struct{}

// NewWalk constructor
func NewWalk() *Walk {
	return &Walk{}
}

// NewRoad constructor
func NewRoad() *Road {
	return &Road{}
}

// BuildRoute method
func (r *Road) BuildRoute(lat float32, lon float32) {
	fmt.Printf("Building Road strategy for latitude %f longitude %f\n", lat, lon)
}

// BuildRoute method
func (w *Walk) BuildRoute(lat float32, lon float32) {
	fmt.Printf("Building Walk strategy for latitude %f longitude %f\n", lat, lon)
}

package builder

import "fmt"

// Car the car class
type Car struct {
	BrandName string
	_Product  Product
}

// NewCar constructor
func NewCar(brandName string) *Car {
	return &Car{BrandName: brandName}
}

// StartUpOperations initial phase
func (c *Car) StartUpOperations() {
	c._Product.Add(fmt.Sprintf("Car model is %s", c.BrandName))
}

// BuildBody create body
func (c *Car) BuildBody() {
	c._Product.Add("this is the body of car")
}

// InsertWheels adding wheels
func (c *Car) InsertWheels() {
	c._Product.Add("4 wheels addon")
}

// AddHeadlights adding header lights
func (c *Car) AddHeadlights() {
	c._Product.Add("2 headlight addon")
}

// EndOperations end of building car
func (c *Car) EndOperations() {
}

// GetVehicle to get the product
func (c *Car) GetVehicle() *Product {
	return &c._Product
}

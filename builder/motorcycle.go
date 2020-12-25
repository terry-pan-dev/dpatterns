package builder

import "fmt"

// Motorcycle class
type Motorcycle struct {
	BrandName string
	_Product  Product
}

// NewMotorcycle constructor
func NewMotorcycle(brandName string) *Motorcycle {
	return &Motorcycle{BrandName: brandName}
}

// StartUpOperations first one
func (m *Motorcycle) StartUpOperations() {
}

// BuildBody essential func
func (m *Motorcycle) BuildBody() {
	m._Product.Add("this is the body of motorcycle")
}

// InsertWheels essential func
func (m *Motorcycle) InsertWheels() {
	m._Product.Add("2 wheels addon")
}

// AddHeadlights essential func
func (m *Motorcycle) AddHeadlights() {
	m._Product.Add("1 headlight addon")
}

// EndOperations essential func
func (m *Motorcycle) EndOperations() {
	m._Product.Add(fmt.Sprintf("Motorcycle model is %s", m.BrandName))
}

// GetVehicle essential func
func (m *Motorcycle) GetVehicle() *Product {
	return &m._Product
}

package builder

import "fmt"

// Product the final vehicle product
type Product struct {
	parts []string
}

// Add add part to product
func (p *Product) Add(part string) {
	p.parts = append(p.parts, part)
}

// ShowProduct print product info
func (p *Product) ShowProduct() {
	for _, part := range p.parts {
		fmt.Println(part)
	}
}

// IVehicleBuilder the abstraction layer of vehicle builder
type IVehicleBuilder interface {
	StartUpOperations()
	BuildBody()
	InsertWheels()
	AddHeadlights()
	EndOperations()
	GetVehicle() *Product
}

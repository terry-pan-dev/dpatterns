package builder

// Director knows how to compose different vehicles
type Director struct {
	_builder IVehicleBuilder
}

// Constructor all the vehicle phases
func (d *Director) Constructor(_builder IVehicleBuilder) {
	d._builder = _builder
	d._builder.StartUpOperations()
	d._builder.BuildBody()
	d._builder.InsertWheels()
	d._builder.AddHeadlights()
	d._builder.EndOperations()
}

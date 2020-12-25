package builder

// MyBuilderAppDemo to demo how to use buidler pattern
func MyBuilderAppDemo() {
	director := Director{}
	var fordCar IVehicleBuilder = NewCar("Ford")
	director.Constructor(fordCar)
	p1 := fordCar.GetVehicle()
	p1.ShowProduct()
	var hondaMotorcycle IVehicleBuilder = NewMotorcycle("Honda")
	director.Constructor(hondaMotorcycle)
	p2 := hondaMotorcycle.GetVehicle()
	p2.ShowProduct()
}

package factory

// AnimalFactory interface
type AnimalFactory interface {
	CreateAnimal()
}

// DogFactory class
type DogFactory struct{}

// CatFactory class
type CatFactory struct{}

// CreateAnimal method
func (d *DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

// CreateAnimal method
func (c *CatFactory) CreateAnimal() Animal {
	return &Cat{}
}

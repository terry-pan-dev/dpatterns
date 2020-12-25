package factory

import "fmt"

// MyFactoryShowApp method
func MyFactoryShowApp() {
	fmt.Println("Factory Design Pattern Demo")
	dogFactory := DogFactory{}
	var dog Animal = dogFactory.CreateAnimal()
	dog.Speak()
	dog.PreferedAction()

	catFatory := CatFactory{}
	var cat Animal = catFatory.CreateAnimal()
	cat.Speak()
	cat.PreferedAction()
}

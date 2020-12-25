package factory

import "fmt"

// MyFactoryAppDemo method
func MyFactoryAppDemo() {
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

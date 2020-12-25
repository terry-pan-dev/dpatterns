package main

import (
	"fmt"

	"github.com/terry-pan-dev/design-pattern/builder"
	"github.com/terry-pan-dev/design-pattern/factory"
)

func main() {
	fmt.Println("============== Factory Pattern ===============")
	factory.MyFactoryShowApp()
	fmt.Println("============== Builder Pattern ===============")
	builder.MyBuilderAppDemo()
}

package main

import (
	"fmt"

	"github.com/terry-pan-dev/design-pattern/abfactory"
	"github.com/terry-pan-dev/design-pattern/bridge"
	"github.com/terry-pan-dev/design-pattern/builder"
	"github.com/terry-pan-dev/design-pattern/factory"
	"github.com/terry-pan-dev/design-pattern/observer"
	"github.com/terry-pan-dev/design-pattern/strategy"
)

func main() {
	fmt.Println("============== Factory Pattern ===============")
	factory.MyFactoryAppDemo()
	fmt.Println("============== Builder Pattern ===============")
	builder.MyBuilderAppDemo()
	fmt.Println("============== Observer Pattern ===============")
	observer.MyObserverAppDemo()
	fmt.Println("============== Strategy Pattern ===============")
	strategy.MyStrategyAppDemo()
	fmt.Println("============== abstract factory Pattern ===============")
	abfactory.MyAbstractFactoryAppDemo()
	fmt.Println("============== brdige Pattern ===============")
	bridge.MyBridgeAppDemo()
}

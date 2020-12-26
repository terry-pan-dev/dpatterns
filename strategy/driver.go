package strategy

// MyStrategyAppDemo driver function
func MyStrategyAppDemo() {
	con := NewContext()
	var walk Strategy = NewWalk()
	con.SetStrategy(walk)
	con.BuildRoute(109.11, 120.12)
	var road Strategy = NewRoad()
	con.SetStrategy(road)
	con.BuildRoute(109.11, 120.12)
}

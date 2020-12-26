package strategy

// Context strategy context
type Context struct {
	_Strategy Strategy
}

// NewContext constructor
func NewContext() *Context {
	return &Context{}
}

// SetStrategy method
func (c *Context) SetStrategy(strategy Strategy) {
	c._Strategy = strategy
}

// BuildRoute method
func (c *Context) BuildRoute(lat float32, lon float32) {
	c._Strategy.BuildRoute(lat, lon)
}

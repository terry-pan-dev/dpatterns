package bridge

// Device interface
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int8
	SetVolume(int8)
	GetChannel() int16
	SetChannel(int16)
}

// Radio class
type Radio struct {
	volume  int8
	channel int16
	enabled bool
}

// TV class
type TV struct {
	volume  int8
	channel int16
	enabled bool
}

// NewRadio constructor
func NewRadio() *Radio {
	return &Radio{}
}

// NewTV constructor
func NewTV() *TV {
	return &TV{}
}

// IsEnabled method
func (r *Radio) IsEnabled() bool {
	return r.enabled
}

// Enable method
func (r *Radio) Enable() {
	r.enabled = true
}

// Disable method
func (r *Radio) Disable() {
	r.enabled = false
}

// GetVolume method
func (r *Radio) GetVolume() int8 {
	return r.volume
}

// SetVolume method
func (r *Radio) SetVolume(volume int8) {
	r.volume = volume
}

// GetChannel method
func (r *Radio) GetChannel() int16 {
	return r.channel
}

// SetChannel method
func (r *Radio) SetChannel(channel int16) {
	r.channel = channel
}

// IsEnabled method
func (t *TV) IsEnabled() bool {
	return t.enabled
}

// Enable method
func (t *TV) Enable() {
	t.enabled = true
}

// Disable method
func (t *TV) Disable() {
	t.enabled = false
}

// GetVolume method
func (t *TV) GetVolume() int8 {
	return t.volume
}

// SetVolume method
func (t *TV) SetVolume(volume int8) {
	t.volume = volume
}

// GetChannel method
func (t *TV) GetChannel() int16 {
	return t.channel
}

// SetChannel method
func (t *TV) SetChannel(channel int16) {
	t.channel = channel
}

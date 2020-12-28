package bridge

const (
	VOL_CHANGE_RATE = 10
	MAX_VOL         = 100
	MIN_VOL         = 0
)

// RemoteControl class
type RemoteControl struct {
	device Device
}

// NewRemoteControl constructor
func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

// TogglePower method
func (r *RemoteControl) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

// VolumeDown method
func (r *RemoteControl) VolumeDown() {
	if r.device.GetVolume()-VOL_CHANGE_RATE <= MIN_VOL {
		r.device.SetVolume(MIN_VOL)
	} else {
		r.device.SetVolume(r.device.GetVolume() - VOL_CHANGE_RATE)
	}
}

// VolumeUp method
func (r *RemoteControl) VolumeUp() {
	if r.device.GetVolume()+VOL_CHANGE_RATE >= MAX_VOL {
		r.device.SetVolume(MAX_VOL)
	} else {
		r.device.SetVolume(r.device.GetVolume() + VOL_CHANGE_RATE)
	}
}

// ChannelDown method
func (r *RemoteControl) ChannelDown() {
	r.device.SetChannel(r.device.GetChannel() - 1)
}

// ChannelUp method
func (r *RemoteControl) ChannelUp() {
	r.device.SetChannel(r.device.GetChannel() + 1)
}

// AdvancedRemoteControl class
type AdvancedRemoteControl struct {
	RemoteControl
}

// NewAdvancedRemoteControl constructor
func NewAdvancedRemoteControl(device Device) *AdvancedRemoteControl {
	return &AdvancedRemoteControl{RemoteControl{device: device}}
}

// Mute method
func (a *AdvancedRemoteControl) Mute() {
	a.device.SetVolume(0)
}

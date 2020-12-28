package bridge

import "fmt"

// MyBridgeAppDemo demo method
func MyBridgeAppDemo() {
	var tv Device = NewTV()
	tvRemote := NewRemoteControl(tv)
	fmt.Printf("Before toggle the state is %v\n", tv.IsEnabled())
	fmt.Println("Toggle power button")
	tvRemote.TogglePower()
	fmt.Printf("Now the state is %v\n", tv.IsEnabled())
	fmt.Printf("Before volumeup the state is %v\n", tv.GetVolume())
	fmt.Println("Volume Up")
	tvRemote.VolumeUp()
	fmt.Printf("After volumeup the state is %v\n", tv.GetVolume())

	var radio Device = NewRadio()
	radioRemote := NewAdvancedRemoteControl(radio)
	fmt.Println("Increse volume")
	radioRemote.VolumeUp()
	fmt.Printf("Now the volume is %d\n", radio.GetVolume())
	fmt.Println("Mute radio volume")
	radioRemote.Mute()
	fmt.Printf("Now the volume is %d\n", radio.GetVolume())
}

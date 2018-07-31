package roku

import (
	"fmt"
	"net/http"
)

func (r Roku) pressKey(key string) error {
	_, err := http.Post(fmt.Sprintf("%skeypress/%s", r.APIAddress, key), "application/json", nil)
	return err
}

// Press the home button for a Roku
func (r Roku) PressHome() error {
	return r.pressKey("Home")
}

// Press the reverse button for a Roku
func (r Roku) PressReverse() error {
	return r.pressKey("Rev")
}

// Press the forward button for a Roku
func (r Roku) PressForward() error {
	return r.pressKey("Fwd")
}

// Press the play button for a Roku
func (r Roku) PressPlay() error {
	return r.pressKey("Play")
}

// Press the select button for a Roku
func (r Roku) PressSelect() error {
	return r.pressKey("Select")
}

// Press the left button for a Roku
func (r Roku) PressLeft() error {
	return r.pressKey("Left")
}

// Press the right button for a Roku
func (r Roku) PressRight() error {
	return r.pressKey("Right")
}

// Press the down button for a Roku
func (r Roku) PressDown() error {
	return r.pressKey("Down")
}

// Press the up button for a Roku
func (r Roku) PressUp() error {
	return r.pressKey("Play")
}

// Press the back button for a Roku
func (r Roku) PressBack() error {
	return r.pressKey("Back")
}

// Press the instant replay button for a Roku
func (r Roku) PressInstantReplay() error {
	return r.pressKey("InstantReplay")
}

// Press the info button for a Roku
func (r Roku) PressInfo() error {
	return r.pressKey("Info")
}

// Press the backspace button for a Roku
func (r Roku) PressBackspace() error {
	return r.pressKey("Backspace")
}

// Press the search button for a Roku
func (r Roku) PressSearch() error {
	return r.pressKey("Search")
}

// Press the enter button for a Roku
func (r Roku) PressEnter() error {
	return r.pressKey("Enter")
}

// Press the power off button for a Roku
func (r Roku) PressPowerOff() error {
	return r.pressKey("PowerOff")
}

// Press the power button for a Roku
func (r Roku) PressPower() error {
	return r.pressKey("Power")
}

// Press the volume up button for a Roku
func (r Roku) PressVolumeUp() error {
	return r.pressKey("VolumeUp")
}

// Press the volume down button for a Roku
func (r Roku) PressVolumeDown() error {
	return r.pressKey("VolumeDown")
}

// PressVolumeMute simulates a mute button press
func (r Roku) PressVolumeMute() error {
	return r.pressKey("VolumeMute")
}

// Press the channel up button for a Roku
func (r Roku) PressChannelUp() error {
	return r.pressKey("ChannelUp")
}

// Press the channel down button for a Roku
func (r Roku) PressChannelDown() error {
	return r.pressKey("ChannelDown")
}

// PressInputTuner simulates an "input tuner" button press
func (r Roku) PressInputTuner() error {
	return r.pressKey("InputTuner")
}

// PressInputHDMI1 simulates an input HDMI1 button press
func (r Roku) PressInputHDMI1() error {
	return r.pressKey("InputHDMI1")
}

// PressInputHDMI2 simulates an input HDMI2 button press
func (r Roku) PressInputHDMI2() error {
	return r.pressKey("InputHDMI2")

}

// PressInputHDMI3 simulates an input HDMI3 button press
func (r Roku) PressInputHDMI3() error {
	return r.pressKey("InputHDMI3")

}

// PressInputHDMI4 simulates an input HDMI4 button press
func (r Roku) PressInputHDMI4() error {
	return r.pressKey("InputHDMI4")

}

// PressInputAVI simulates an input AVI button press
func (r Roku) PressInputAVI() error {
	return r.pressKey("InputAVI")
}

// TODO: character support

/*
File to provide the types for the gorecon package.
Also contains some simple helper functions for the types, e.g. the String functions
*/

package reconlib

import "context"
import "fmt"
import "github.com/google/gousb"

// FanController is an instance of the bitfenix recon fan controller device, connect by USB
type FanController struct {
	context                      *gousb.Context
	device                       *gousb.Device
	deviceInterface              *gousb.Interface
	deviceInterfaceCloseFunction func()
	communicationEndpoint        *gousb.InEndpoint
	readStream                   *gousb.ReadStream
	readContext                  context.Context
	readCancelFunc               context.CancelFunc
	waitReadDone                 chan struct{}
}

// ------------------------------------

// Temperature struct to give temperature measurements in celsius and fahrenheit.
type Temperature struct {
	Celsius    uint8
	Fahrenheit uint8
}

func (temperature Temperature) String() string {
	return fmt.Sprintf("%d°C (%d°F)", temperature.Celsius, temperature.Fahrenheit)
}

// ------------------------------------

// Speed represents the speed of a fan in rounds per minute (RPM).
type Speed uint16

func (speed Speed) String() string {
	return fmt.Sprintf("%d RPM", speed)
}

// ------------------------------------

// Channel is the channel number of the controller. Ranges from 1 to 5.
type Channel uint8

const (
	// Channel1 is the first channel of the fan controller.
	Channel1 Channel = 0
	// Channel2 is the second channel of the fan controller.
	Channel2 Channel = 1
	// Channel3 is the third channel of the fan controller.
	Channel3 Channel = 2
	// Channel4 is the fourth channel of the fan controller.
	Channel4 Channel = 3
	// Channel5 is the fifth channel of the fan controller.
	Channel5 Channel = 4
)

func (channel Channel) String() string {
	return fmt.Sprintf("Channel %d", channel+1)
}

// ------------------------------------

// Mode is the device mode. It can be manual (ManualMode) or automatic (AutomaticMode).
type Mode uint8

const (
	// ManualMode is used to control the fan speed manually.
	ManualMode Mode = 0
	// AutomaticMode is used to let the controller manage the fan speed according to current temperature readings.
	AutomaticMode Mode = 1
)

func (mode Mode) String() string {
	if mode == ManualMode {
		return "Manual Mode"
	}
	return "Automatic Mode"
}

// ------------------------------------

// Display is the status of the LCD Display. It can either be on or off.
type Display uint8

const (
	// DisplayOff indicates that the display is off.
	DisplayOff Display = 0
	// DisplayOn indicates that the display is on.
	DisplayOn Display = 1
)

func (display Display) String() string {
	if display == DisplayOff {
		return "Display Off"
	}
	return "Display On"
}

// ------------------------------------

// Sound indicates the sound setting of the controller (beeping noise on touch).
type Sound uint8

const (
	// SoundOff the sound of the device is disabled.
	SoundOff Sound = 0
	// SoundOn the sound of the device is enabled.
	SoundOn Sound = 1
)

func (sound Sound) String() string {
	if sound == SoundOff {
		return "Sound Off"
	}
	return "Sound On"
}

// ------------------------------------

// Settings of the fan controller.
type Settings struct {
	Sound   Sound
	Display Display
	Channel Channel
	Mode    Mode
}

func (settings Settings) String() string {
	return fmt.Sprintf("Settings: %s, %s, %s, Active %s",
		settings.Sound, settings.Display, settings.Mode, settings.Channel)
}

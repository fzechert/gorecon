package reconlib

import "fmt"

// flag to enable debug messages
var debug = false
var usbdebug = false

// EnableDebug enables or disables debug output (to stdout).
func EnableDebug(enabled bool) {
	debug = enabled
}

// EnableUsbDebug enables or disables usb debug output (to stdout).
func EnableUsbDebug(enabled bool) {
	usbdebug = enabled
}

func debugMessage(message string) {
	if debug {
		fmt.Printf("[debug] %v\n", message)
	}
}

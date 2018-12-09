package main

import "flag"
import "fmt"
import "github.com/fzechert/gorecon/pkg/reconlib"
import "time"

func main() {
	enableDebugPtr := flag.Bool("debug", false, "enable debug messages")
	enableUsbDebugPtr := flag.Bool("debugusb", false, "enable usb debug messages")

	flag.Parse()

	reconlib.EnableDebug(*enableDebugPtr)
	reconlib.EnableUsbDebug(*enableUsbDebugPtr)

	controller, error := reconlib.Connect()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("%v\n", controller)
	time.Sleep(time.Duration(30) * time.Second)
	controller.Close()
}

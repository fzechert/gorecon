/*
This file contains the main functionality of the gorecon pacakge.
E.g. it contains all the methods that can be used to interact with the actual
bitfenix recon fan controller.
*/

package reconlib

import "context"
import "errors"
import "fmt"
import "github.com/google/gousb"

/*
Connect to the bit fenix recon fan controller.

Use this function only if there is only one controller connected via USB.
If you have multiple controller then use the Connect(int) function to choose
the correct controller you want to connect to. Returns an error if the connection
cannot be made because there is no controller or there are more than one.
This is equivalent to calling ConnectTo(0).
*/
func Connect() (*FanController, error) {
	return ConnectTo(0)
}

/*
ConnectTo connects to the bit fenix recon fan controller at index deviceIndex.

Use this function to connect to the recon fan controller with the given device index.
The device index starts at 0 and increases by 1 for every fan controller that is connected
to your USB port. If you have two fan controllers connect to USB use the function
Connect(0) to connect to the first one and Connect(1) to connect to the second one.
The function will return an error if the connection to the specified controller cannot be made.
*/
func ConnectTo(deviceIndex int) (*FanController, error) {
	controller := new(FanController)
	controller.context = gousb.NewContext()

	if usbdebug {
		controller.context.Debug(libusbLogDebug)
	} else {
		controller.context.Debug(libusbLogNone)
	}

	vid, pid, openIndex := gousb.ID(vendorID), gousb.ID(productID), 0
	debugMessage(fmt.Sprintf("trying to connect to %v:%v at index %d", vid, pid, deviceIndex))

	devices, err := controller.context.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if desc.Vendor == vid && desc.Product == pid {
			openIndex++
			if openIndex == deviceIndex+1 {
				debugMessage(fmt.Sprintf("found and selected usb device %v:%v at index %d",
					desc.Vendor, desc.Product, openIndex-1))
				return true
			}
			debugMessage(fmt.Sprintf("found usb device %v:%v at wrong index %d",
				desc.Vendor, desc.Product, openIndex-1))
		} else {
			debugMessage(fmt.Sprintf("found other usb device %v:%v", desc.Vendor, desc.Product))
		}
		return false
	})

	// an error occurred during opening of the device
	if err != nil {
		// there are no devices
		if len(devices) == 0 {
			defer controller.context.Close()
			debugMessage(fmt.Sprintf("no recon fan controller found or not possible to open: %v", err))
			return nil, fmt.Errorf("no recon fan controller found or not possible to open: %v", err)
		}

		defer controller.context.Close()
		for _, device := range devices {
			defer device.Close()
		}

		debugMessage(fmt.Sprintf("an error occurred accessing your usb devices: %v", err))
		return nil, fmt.Errorf("an error occurred accessing your usb devices: %v", err)
	}

	controller.device = devices[0]

	debugMessage("connected to device, device info:")
	debugMessage(fmt.Sprintf("	device %v:%v",
		controller.device.Desc.Vendor, controller.device.Desc.Product))
	debugMessage(fmt.Sprintf("	address %d:%d:%d",
		controller.device.Desc.Bus, controller.device.Desc.Address, controller.device.Desc.Port))
	debugMessage(fmt.Sprintf("	USB version: %v, device version: %v",
		controller.device.Desc.Spec, controller.device.Desc.Device))
	debugMessage(fmt.Sprintf("	device speed: %v", controller.device.Desc.Speed))
	debugMessage(fmt.Sprintf("	device class: %v:%v:%v",
		controller.device.Desc.Class, controller.device.Desc.SubClass, controller.device.Desc.Protocol))
	debugMessage(fmt.Sprintf("	max control package size: %v", controller.device.Desc.MaxControlPacketSize))

	// make sure the usb device is automatically detached from the kernel when we want to access it
	// otherwise we will not gain access to the device as it is exclusively held by the kernel
	controller.device.SetAutoDetach(true)

	// connect to the default interface
	devInterface, devInterfaceClose, err := controller.device.DefaultInterface()
	if err != nil {
		debugMessage(fmt.Sprintf("failed to select default interface: %v", err))
		defer controller.context.Close()
		defer controller.device.Close()
		if devInterfaceClose != nil {
			defer devInterfaceClose()
		}

		return nil, fmt.Errorf("failed to select default interface: %v", err)
	}

	controller.deviceInterface = devInterface
	controller.deviceInterfaceCloseFunction = devInterfaceClose

	// find the communication endpoint
	for _, endpointDescription := range controller.deviceInterface.Setting.Endpoints {
		if endpointDescription.Direction == gousb.EndpointDirectionIn {
			debugMessage(fmt.Sprintf("found endpoint for direction in: %v", endpointDescription))
			endpoint, e := controller.deviceInterface.InEndpoint(endpointDescription.Number)
			if e != nil {
				debugMessage(fmt.Sprintf("failed to select the endpoint %v: %v", endpointDescription, e))
				defer controller.context.Close()
				defer controller.device.Close()
				defer devInterfaceClose()
				return nil, fmt.Errorf("failed to select the endpoint %v: %v", endpointDescription, e)
			}
			debugMessage(fmt.Sprintf("selected endpoint for direction in: %v", endpointDescription))
			controller.communicationEndpoint = endpoint
			break
		}
	}

	if controller.communicationEndpoint == nil {
		debugMessage("found no suitable communication endpoint")
		defer controller.context.Close()
		defer controller.device.Close()
		defer devInterfaceClose()
		return nil, errors.New("found no suitable communication endpoint")
	}

	debugMessage("successfully connected to recon fan controller")

	reader, err := controller.communicationEndpoint.NewStream(
		controller.communicationEndpoint.Desc.MaxPacketSize, gousbParallelRead)

	if err != nil {
		debugMessage(fmt.Sprintf("failed to open the endpoint for reading: %v", err))
		defer controller.context.Close()
		defer controller.device.Close()
		defer devInterfaceClose()
		return nil, fmt.Errorf("failed to open the endpoint for reading: %v", err)
	}

	controller.readStream = reader
	ctx, ctxCancel := context.WithCancel(context.Background())
	controller.readContext = ctx
	controller.readCancelFunc = ctxCancel
	controller.waitReadDone = make(chan struct{})
	go controller.read()
	return controller, nil
}

// Close the fan controller
// After closing the fan controller you should not use the fan controller instance any longer.
func (controller *FanController) Close() {
	if controller.context != nil {
		defer controller.context.Close()
	}
	if controller.device != nil {
		defer controller.device.Close()
	}
	if controller.deviceInterfaceCloseFunction != nil {
		defer controller.deviceInterfaceCloseFunction()
	}
	if controller.readContext != nil {
		controller.readCancelFunc()
		debugMessage("shutdown: waiting for the read operation to finish")
		<-controller.waitReadDone
		debugMessage("shutdown: read completed")
	}
}

func (controller *FanController) read() {
	for {
		select {
		case <-controller.readContext.Done():
			close(controller.waitReadDone)
			return
		default:
			buffer := make([]byte, controller.communicationEndpoint.Desc.MaxPacketSize)
			read, err := controller.readStream.ReadContext(controller.readContext, buffer)
			debugMessage(fmt.Sprintf("read %d bytes with error: %v >>%v", read, err, buffer))
			if err != nil {
				debugMessage(fmt.Sprintf("error during read: %v", err))
				continue
			}

			p, err := parsePacket(buffer)
			debugMessage(fmt.Sprintf("parsed packet: %v with error %v", p, err))
		}
	}
}

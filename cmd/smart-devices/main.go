package main

import (
	"fmt"

	"github.com/dscholtz/interface-basics.git/pkg/device"
)

func UseDevice(d device.SmartDevice) {
	d.Charge()
	d.PlayMusic()
	d.TakePhoto()
}

func main() {
	iphone := device.Phone{Brand: "iPhone"}
	samsung := device.Phone{Brand: "Samsung"}

	fmt.Println("Using iPhone:")
	UseDevice(iphone)

	fmt.Println("\nUsing Samsung")
	UseDevice(samsung)

	// cannot do the following becuase tablet does not satisfy the
	// composed interface for SmartDevice.
	// tablet := device.Tablet{Brand: "Lenovo"}
	// UseDevice(tablet)
}

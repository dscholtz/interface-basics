package device

import "fmt"

// implement a phone device with all feature interfaces
type Phone struct {
	Brand string
}

func (p Phone) Charge() {
	fmt.Println(p.Brand, "is charging")
}

func (p Phone) PlayMusic() {
	fmt.Println(p.Brand, "is playing music")
}

func (p Phone) TakePhoto() {
	fmt.Println(p.Brand, "took a photo")
}

// this will automatically make the Phone qualify as a SmartDevice
// becuase it has all interfaces that is needed for that composed interface.

// lets make a tablet with partial features, no camera
type Tablet struct {
	Brand string
}

func (t Tablet) Charge() {
	fmt.Println(t.Brand, "is charging")
}

func (t Tablet) PlayMusic() {
	fmt.Println(t.Brand, "is playing music")
}

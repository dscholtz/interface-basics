package device

// lets create small interfaces for each device
// each interface describes one feature
type Charger interface {
	Charge()
}

type MusicPlayer interface {
	PlayMusic()
}

type Camera interface {
	TakePhoto()
}

// lets compose these interfaces into a composed interface.
// --> full featured smart device.
type SmartDevice interface {
	Charger
	MusicPlayer
	Camera
}

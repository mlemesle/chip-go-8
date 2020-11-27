package beeper

// Mute is a mute beeper, meaning no sound is produced
type Mute struct {
}

// NewMute creates a new instance of a mute beeper
func NewMute() *Mute {
	return &Mute{}
}

// Init initializes the given Mute
func (b *Mute) Init() error {
	return nil
}

// Beep does nothing
func (b *Mute) Beep() {}

// Destroy free SDL allocated components
func (b *Mute) Destroy() {}

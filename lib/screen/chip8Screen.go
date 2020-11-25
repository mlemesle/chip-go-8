package screen

import "github.com/mlemesle/chip-go-8/lib/emulator"

// Chip8ScreenInterface represents the methods that needs to be implemented by the screen
type Chip8ScreenInterface interface {
	Init() error
	Destroy()
	Draw(c *emulator.Chip8) error
	HandleEvent(c *emulator.Chip8) bool
}

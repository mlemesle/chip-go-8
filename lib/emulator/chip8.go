package emulator

import (
	"errors"
	"os"
)

const (
	memorySize    = 4096
	memoryOffset  = 512
	registersSize = 16
	gfxSize       = 32 * 64
	stackSize     = 16
	keySize       = 16
	fontSetSize   = 80
)

// The Chip8's font set
var chip8FontSet = [fontSetSize]uint16{
	0xF0, 0x90, 0x90, 0x90, 0xF0, //0
	0x20, 0x60, 0x20, 0x20, 0x70, //1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
	0x90, 0x90, 0xF0, 0x10, 0x10, //4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
	0xF0, 0x10, 0x20, 0x40, 0x40, //7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
	0xF0, 0x90, 0xF0, 0x90, 0x90, //A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
	0xF0, 0x80, 0x80, 0x80, 0xF0, //C
	0xE0, 0x90, 0x90, 0x90, 0xE0, //D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
	0xF0, 0x80, 0xF0, 0x80, 0x80, //F
}

// Chip8 is the representation of a chip 8 emulator (https://fr.wikipedia.org/wiki/CHIP-8)
type Chip8 struct {
	opcode     uint16
	memory     [memorySize]uint16
	registers  [registersSize]uint8
	i          uint16
	pc         uint16
	gfx        [gfxSize]uint8
	delayTimer uint8
	soundTimer uint8
	stack      [stackSize]uint16
	sp         byte
	key        [keySize]byte
	draw       bool
}

// Chip8Interface is the set of method the emulator needs to implement
type Chip8Interface interface {
	Initialize()
	NeedDraw() bool
	SetDraw(b bool)
	GetGFX() [gfxSize]uint8
	SetRegisterUp(index int)
	SetRegisterDown(index int)
	LoadMemory(filename string) error
	EmulateCycle(filename string) error
}

// New creates a default and non-initialized emulator
func New() *Chip8 {
	return &Chip8{}
}

// Initialize sets defaults value to all fields of the emulator
func (c *Chip8) Initialize() {
	c.opcode = 0
	c.memory = [memorySize]uint16{}
	for i := 0; i < fontSetSize; i++ {
		c.memory[i] = chip8FontSet[i]
	}
	c.registers = [registersSize]uint8{}
	c.i = 0
	c.pc = 0x200
	c.gfx = [gfxSize]uint8{}
	c.delayTimer = 0
	c.soundTimer = 0
	c.stack = [stackSize]uint16{}
	c.sp = 0
	c.key = [keySize]byte{}
	c.draw = false
}

// NeedDraw tells if the emulator needs to draw on the display
func (c Chip8) NeedDraw() bool {
	return c.draw
}

// SetDraw sets the value of the 'draw' field
func (c *Chip8) SetDraw(b bool) {
	c.draw = b
}

// GetGFX gets the gfx of the emulator
func (c Chip8) GetGFX() [gfxSize]uint8 {
	return c.gfx
}

// SetKeyUp sets the value to 'up' for the given key index
func (c *Chip8) SetKeyUp(index int) {
	c.key[index] = 0
}

// SetKeyDown sets the value to 'down' for the given key index
func (c *Chip8) SetKeyDown(index int) {
	c.key[index] = 1
}

// LoadMemory load the file in parameter into the emulator's memory
func (c *Chip8) LoadMemory(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Size() > memorySize-0x200 {
		return errors.New("File is too big for memory")
	}

	fileBuffer := make([]byte, stat.Size())
	if _, err = file.Read(fileBuffer); err != nil {
		return err
	}

	bufferLen := len(fileBuffer)
	for i := 0; i < bufferLen; i++ {
		c.memory[i+memoryOffset] = uint16(fileBuffer[i])
	}
	return nil
}

// EmulateCycle emulate a cycle of the emulator's processor
func (c *Chip8) EmulateCycle() error {
	c.opcode = uint16(c.memory[c.pc]<<8) | uint16(c.memory[c.pc+1])

	err := handleOpcode(c)
	if err != nil {
		return err
	}

	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		if c.soundTimer == 1 {
			println("beep")
		}
		c.soundTimer--
	}

	return nil
}

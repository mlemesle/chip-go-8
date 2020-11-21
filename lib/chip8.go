package lib

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

type Chip8 struct {
	opcode     uint16
	memory     [memorySize]uint16
	registers  [registersSize]uint16
	i          uint16
	pc         uint16
	gfx        [gfxSize]uint8
	delayTimer uint16
	soundTimer uint16
	stack      [stackSize]uint16
	sp         byte
	key        [keySize]byte
	draw       bool
	width      int32
	height     int32
}

type Chip8Emulator interface {
	Initialize()
	NeedDraw() bool
	GetGFX() [gfxSize]uint8
	SetRegisterUp(index int)
	SetRegisterDown(index int)
	LoadMemory(filename string) error
	EmulateCycle(filename string) error
}

func (c *Chip8) Initialize() {
	c.opcode = 0
	c.memory = [memorySize]uint16{}
	for i := 0; i < fontSetSize; i++ {
		c.memory[i] = chip8FontSet[i]
	}
	c.registers = [registersSize]uint16{}
	c.i = 0
	c.pc = 0x200
	c.gfx = [gfxSize]uint8{}
	c.delayTimer = 0
	c.soundTimer = 0
	c.stack = [stackSize]uint16{}
	c.sp = 15
	c.key = [keySize]byte{}
	c.draw = false
	println("Chip8 init")
}

func (c Chip8) NeedDraw() bool {
	return c.draw
}

func (c Chip8) GetGFX() [gfxSize]uint8 {
	return c.gfx
}

func (c *Chip8) SetKeyUp(index int) {
	c.key[index] = 0
}

func (c *Chip8) SetKeyDown(index int) {
	c.key[index] = 1
}

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
	println("Memory loaded")
	return nil
}

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

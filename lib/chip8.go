package lib

import (
	"errors"
	"os"
)

const (
	memorySize    = 4096
	memoryOffset  = 512
	registersSize = 16
	gfxXSize      = 32
	gfxYSize      = 64
	stackSize     = 16
	keySize       = 16
	fontSetSize   = 80
)

var chip8FontSet = [fontSetSize]uint8{
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
	memory     [memorySize]byte
	registers  [registersSize]byte
	i          uint16
	pc         uint16
	gfx        [gfxXSize][gfxYSize]uint8
	delayTimer byte
	soundTimer byte
	stack      [stackSize]uint16
	sp         byte
	key        [keySize]byte
	draw       bool
}

type Chip8Emulator interface {
	Initialize()
	LoadMemory(filename string) error
	EmulateCycle(filename string) error
}

func (c *Chip8) Initialize() {
	c.opcode = 0
	c.memory = [memorySize]byte{}
	for i := 0; i < fontSetSize; i++ {
		c.memory[i] = chip8FontSet[i]
	}
	c.registers = [registersSize]byte{}
	c.i = 0
	c.pc = 0x200
	c.gfx = [gfxXSize][gfxYSize]uint8{}
	c.delayTimer = 0
	c.soundTimer = 0
	c.stack = [stackSize]uint16{}
	c.sp = 0
	c.key = [keySize]byte{}
	c.draw = false
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
	if stat.Size() > memorySize-512 {
		return errors.New("File is too big for memory")
	}

	fileBuffer := make([]byte, stat.Size())
	if _, err = file.Read(fileBuffer); err != nil {
		return err
	}

	bufferLen := len(fileBuffer)
	for i := 0; i < bufferLen; i++ {
		c.memory[i+memoryOffset] = fileBuffer[i]
	}
	return nil
}

func (c *Chip8) EmulateCycle() error {
	c.opcode = uint16(c.memory[c.pc]<<8) | uint16(c.memory[c.pc+1])

	opcodeFn, err := getOpcodeFunc(c.opcode)
	if err != nil {
		return err
	}

	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		if c.soundTimer == 1 {
			// TODO Do beep
		}
		c.soundTimer--
	}

	return nil
}

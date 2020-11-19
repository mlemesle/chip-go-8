package lib

import (
	"errors"
	"math/rand"
)

func getOpcodeFunc(c *Chip8) error {
	if c.opcode == 0x00EE {
		opcode_00EE(c)
	} else if c.opcode == 0x00E0 {
		opcode_00E0(c)
	} else if (c.opcode & 0xF000) == 0x0000 {
		opcode_0NNN(c)
	} else if (c.opcode & 0xF000) == 0x1000 {
		opcode_1NNN(c)
	} else if (c.opcode & 0xF000) == 0x2000 {
		opcode_2NNN(c)
	} else if (c.opcode & 0xF000) == 0x3000 {
		opcode_3XNN(c)
	} else if (c.opcode & 0xF000) == 0x4000 {
		opcode_4XNN(c)
	} else if (c.opcode & 0xF00F) == 0x5000 {
		opcode_5XY0(c)
	} else if (c.opcode & 0xF000) == 0x6000 {
		opcode_6XNN(c)
	} else if (c.opcode & 0xF000) == 0x7000 {
		opcode_7XNN(c)
	} else if (c.opcode & 0xF00F) == 0x8000 {
		opcode_8XY0(c)
	} else if (c.opcode & 0xF00F) == 0x8001 {
		opcode_8XY1(c)
	} else if (c.opcode & 0xF00F) == 0x8002 {
		opcode_8XY2(c)
	} else if (c.opcode & 0xF00F) == 0x8003 {
		opcode_8XY3(c)
	} else if (c.opcode & 0xF00F) == 0x8004 {
		opcode_8XY4(c)
	} else if (c.opcode & 0xF00F) == 0x8005 {
		opcode_8XY5(c)
	} else if (c.opcode & 0xF00F) == 0x8006 {
		opcode_8XY6(c)
	} else if (c.opcode & 0xF00F) == 0x8007 {
		opcode_8XY7(c)
	} else if (c.opcode & 0xF00F) == 0x800E {
		opcode_8XYE(c)
	} else if (c.opcode & 0xF00F) == 0x9000 {
		opcode_9XY0(c)
	} else if (c.opcode & 0xF000) == 0xA000 {
		opcode_ANNN(c)
	} else if (c.opcode & 0xF000) == 0xB000 {
		opcode_BNNN(c)
	} else if (c.opcode & 0xF000) == 0xC000 {
		opcode_CXNN(c)
	} else if (c.opcode & 0xF000) == 0xD000 {
		opcode_DXYN(c)
	} else if (c.opcode & 0xF0FF) == 0xE09E {
		opcode_EX9E(c)
	} else if (c.opcode & 0xF0FF) == 0xE0A1 {
		opcode_EXA1(c)
	} else if (c.opcode & 0xF0FF) == 0xF007 {
		opcode_FX07(c)
	} else if (c.opcode & 0xF0FF) == 0xF00A {
		opcode_FX0A(c)
	} else if (c.opcode & 0xF0FF) == 0xF015 {
		opcode_FX15(c)
	} else if (c.opcode & 0xF0FF) == 0xF018 {
		opcode_FX18(c)
	} else if (c.opcode & 0xF0FF) == 0xF01E {
		opcode_FX1E(c)
	} else if (c.opcode & 0xF0FF) == 0xF029 {
		opcode_FX29(c)
	} else if (c.opcode & 0xF0FF) == 0xF033 {
		opcode_FX33(c)
	} else if (c.opcode & 0xF0FF) == 0xF055 {
		opcode_FX55(c)
	} else if (c.opcode & 0xF0FF) == 0xF065 {
		opcode_FX65(c)
	} else {
		return errors.New("Unknow c.opcode found : " + string(c.opcode))
	}

	return nil
}

func opcode_00EE(c *Chip8) {
	c.sp--
	c.pc = c.stack[c.sp]
	c.pc += 2
}

func opcode_00E0(c *Chip8) {
	c.gfx = [gfxXSize][gfxYSize]uint8{}
	c.draw = true
	c.pc = c.pc + 2
}

func opcode_0NNN(c *Chip8) {
	c.pc += 2
}

func opcode_1NNN(c *Chip8) {
	c.pc = c.opcode & 0x0FFF
}

func opcode_2NNN(c *Chip8) {
	c.stack[c.sp] = c.pc
	c.sp++
	c.pc = c.opcode & 0x0FFF
}

func opcode_3XNN(c *Chip8) {
	var nByteJump uint16 = 2
	if uint16(c.registers[(c.opcode&0x0F00)>>8]) == c.opcode&0x00FF {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_4XNN(c *Chip8) {
	var nByteJump uint16 = 2
	if uint16(c.registers[(c.opcode&0x0F00)>>8]) != c.opcode&0x00FF {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_5XY0(c *Chip8) {
	var nByteJump uint16 = 2
	if c.registers[(c.opcode&0x0F00)>>8] == c.registers[(c.opcode&0x00F0)>>4] {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_6XNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = uint8(c.opcode & 0x00FF)
	c.pc += 2
}

func opcode_7XNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] + uint8(c.opcode&0x00FF)
	c.pc += 2
}

func opcode_8XY0(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY1(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] | c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY2(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] & c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY3(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] ^ c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY4(c *Chip8) {
	var carry byte = 0
	if c.registers[(c.opcode&0x00F0)>>4] > 0xFF-c.registers[(c.opcode&0x0F00)>>8] {
		carry = 1
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] + c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY5(c *Chip8) {
	var carry byte = 1
	if c.registers[(c.opcode&0x00F0)>>4] > c.registers[(c.opcode&0x0F00)>>8] {
		carry = 0
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] - c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

func opcode_8XY6(c *Chip8) {
	c.registers[0xF] = c.registers[(c.opcode&0x0F00)>>8] & 0x1
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] >> 1
	c.pc += 2
}

func opcode_8XY7(c *Chip8) {
	var carry byte = 1
	if c.registers[(c.opcode&0x0F00)>>8] > c.registers[(c.opcode&0x00F0)>>4] {
		carry = 0
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x00F0)>>4] - c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

func opcode_8XYE(c *Chip8) {
	c.registers[0xF] = c.registers[(c.opcode&0x0F00)>>8] >> 7
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] << 1
	c.pc += 2
}

func opcode_9XY0(c *Chip8) {
	var nByteJump uint16 = 2
	if c.registers[(c.opcode&0x0F00)>>8] != c.registers[(c.opcode&0x00F0)>>4] {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_ANNN(c *Chip8) {
	c.i = c.opcode & 0x0FFF
	c.pc += 2
}

func opcode_BNNN(c *Chip8) {
	c.pc = (c.opcode & 0x0FFF) + uint16(c.registers[0x0])
}

func opcode_CXNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = uint8(rand.Intn(256)) & uint8(c.opcode&0x00FF)
	c.pc = c.pc + 2
}

func opcode_DXYN(c *Chip8) {
	x := c.registers[(c.opcode&0x0F00)>>8]
	y := c.registers[(c.opcode&0x00F0)>>4]
	h := c.opcode & 0x000F
	c.registers[0xF] = 0
	var j uint16
	var i uint16
	for j = 0; j < h; j++ {
		pixel := c.memory[c.i+j]
		for i = 0; i < 8; i++ {
			if (pixel & (0x80 >> i)) != 0 {
				if c.gfx[(y + uint8(j))][x+uint8(i)] == 1 {
					c.registers[0xF] = 1
				}
				c.gfx[(y + uint8(j))][x+uint8(i)] ^= 1
			}
		}
	}
	c.draw = true
	c.pc += 2
}

func opcode_EX9E(c *Chip8) {
	var nByteJump uint16 = 2
	if c.key[c.registers[(c.opcode&0x0F00)>>8]] == 1 {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_EXA1(c *Chip8) {
	var nByteJump uint16 = 2
	if c.key[c.registers[(c.opcode&0x0F00)>>8]] == 0 {
		nByteJump += 2
	}
	c.pc += nByteJump
}

func opcode_FX07(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.delayTimer
	c.pc += 2
}

func opcode_FX0A(c *Chip8) {
	isPressed := false
	for i, k := range c.key {
		if k != 0 {
			c.registers[(c.opcode&0x0F00)>>8] = uint8(i)
			isPressed = true
		}
	}
	if !isPressed {
		return
	}
	c.pc += 2
}

func opcode_FX15(c *Chip8) {
	c.delayTimer = c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

func opcode_FX18(c *Chip8) {
	c.soundTimer = c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

func opcode_FX1E(c *Chip8) {
	var carry byte = 0
	if c.i+uint16(c.registers[(c.opcode&0x0F00)>>8]) > 0xFFF {
		carry = 1
	}
	c.registers[0xF] = carry
	c.i = c.i + uint16(c.registers[(c.opcode&0x0F00)>>8])
	c.pc += 2
}

func opcode_FX29(c *Chip8) {
	c.i = uint16(c.registers[(c.opcode&0x0F00)>>8]) * 0x5
	c.pc += 2
}

func opcode_FX33(c *Chip8) {
	c.memory[c.i] = c.registers[(c.opcode&0x0F00)>>8] / 100
	c.memory[c.i+1] = (c.registers[(c.opcode&0x0F00)>>8] / 10) % 10
	c.memory[c.i+2] = (c.registers[(c.opcode&0x0F00)>>8] % 100) / 10
	c.pc += 2
}

func opcode_FX55(c *Chip8) {
	for i := 0; i < int((c.opcode&0x0F00)>>8)+1; i++ {
		c.memory[uint16(i)+c.i] = c.registers[i]
	}
	c.i = ((c.opcode & 0x0F00) >> 8) + 1
	c.pc += 2
}

func opcode_FX65(c *Chip8) {
	for i := 0; i < int((c.opcode&0x0F00)>>8)+1; i++ {
		c.registers[i] = c.memory[c.i+uint16(i)]
	}
	c.i = ((c.opcode & 0x0F00) >> 8) + 1
	c.pc += 2
}

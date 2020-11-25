package emulator

import (
	"errors"
	"github.com/mlemesle/chip-go-8/lib/utils"
	"math/rand"
)

// handleOpcode takes the chip8 and process its current opcode
// chip8's attributes are modified
func handleOpcode(c *Chip8) error {
	if c.opcode == 0x00EE {
		opcode00EE(c)
	} else if c.opcode == 0x00E0 {
		opcode00E0(c)
	} else if (c.opcode & 0xF000) == 0x0000 {
		opcode0NNN(c)
	} else if (c.opcode & 0xF000) == 0x1000 {
		opcode1NNN(c)
	} else if (c.opcode & 0xF000) == 0x2000 {
		opcode2NNN(c)
	} else if (c.opcode & 0xF000) == 0x3000 {
		opcode3XNN(c)
	} else if (c.opcode & 0xF000) == 0x4000 {
		opcode4XNN(c)
	} else if (c.opcode & 0xF00F) == 0x5000 {
		opcode5XY0(c)
	} else if (c.opcode & 0xF000) == 0x6000 {
		opcode6XNN(c)
	} else if (c.opcode & 0xF000) == 0x7000 {
		opcode7XNN(c)
	} else if (c.opcode & 0xF00F) == 0x8000 {
		opcode8XY0(c)
	} else if (c.opcode & 0xF00F) == 0x8001 {
		opcode8XY1(c)
	} else if (c.opcode & 0xF00F) == 0x8002 {
		opcode8XY2(c)
	} else if (c.opcode & 0xF00F) == 0x8003 {
		opcode8XY3(c)
	} else if (c.opcode & 0xF00F) == 0x8004 {
		opcode8XY4(c)
	} else if (c.opcode & 0xF00F) == 0x8005 {
		opcode8XY5(c)
	} else if (c.opcode & 0xF00F) == 0x8006 {
		opcode8XY6(c)
	} else if (c.opcode & 0xF00F) == 0x8007 {
		opcode8XY7(c)
	} else if (c.opcode & 0xF00F) == 0x800E {
		opcode8XYE(c)
	} else if (c.opcode & 0xF00F) == 0x9000 {
		opcode9XY0(c)
	} else if (c.opcode & 0xF000) == 0xA000 {
		opcodeANNN(c)
	} else if (c.opcode & 0xF000) == 0xB000 {
		opcodeBNNN(c)
	} else if (c.opcode & 0xF000) == 0xC000 {
		opcodeCXNN(c)
	} else if (c.opcode & 0xF000) == 0xD000 {
		opcodeDXYN(c)
	} else if (c.opcode & 0xF0FF) == 0xE09E {
		opcodeEX9E(c)
	} else if (c.opcode & 0xF0FF) == 0xE0A1 {
		opcodeEXA1(c)
	} else if (c.opcode & 0xF0FF) == 0xF007 {
		opcodeFX07(c)
	} else if (c.opcode & 0xF0FF) == 0xF00A {
		opcodeFX0A(c)
	} else if (c.opcode & 0xF0FF) == 0xF015 {
		opcodeFX15(c)
	} else if (c.opcode & 0xF0FF) == 0xF018 {
		opcodeFX18(c)
	} else if (c.opcode & 0xF0FF) == 0xF01E {
		opcodeFX1E(c)
	} else if (c.opcode & 0xF0FF) == 0xF029 {
		opcodeFX29(c)
	} else if (c.opcode & 0xF0FF) == 0xF033 {
		opcodeFX33(c)
	} else if (c.opcode & 0xF0FF) == 0xF055 {
		opcodeFX55(c)
	} else if (c.opcode & 0xF0FF) == 0xF065 {
		opcodeFX65(c)
	} else {
		return errors.New("Unknow c.opcode found : " + string(c.opcode))
	}

	return nil
}

// RET
// Return from a subroutine.
//
// The interpreter sets the program counter to the address at the top of the stack,
// then subtracts 1 from the stack pointer.
func opcode00EE(c *Chip8) {
	c.sp--
	c.pc = c.stack[c.sp]
	c.pc += 2
}

// CLS
// Clear the display.
func opcode00E0(c *Chip8) {
	c.gfx = [gfxSize]uint8{}
	c.draw = true
	c.pc = c.pc + 2
}

// SYS addr
// Jump to a machine code routine at nnn.
//
// This instruction is only used on the old computers on which Chip-8 was originally implemented.
// It is ignored by modern interpreters.
func opcode0NNN(c *Chip8) {
	c.pc += 2
}

// JP addr
// Jump to location nnn.
//
// The interpreter sets the program counter to nnn.
func opcode1NNN(c *Chip8) {
	c.pc = c.opcode & 0x0FFF
}

// CALL addr
// Call subroutine at nnn.
//
// The interpreter increments the stack pointer, then puts the current PC on the top of the stack.
// The PC is then set to nnn.
func opcode2NNN(c *Chip8) {
	c.stack[c.sp] = c.pc
	c.sp++
	c.pc = c.opcode & 0x0FFF
}

// SE Vx, byte
// Skip next instruction if Vx = kk.
//
// The interpreter compares register Vx to kk, and if they are equal,
// increments the program counter by 2.
func opcode3XNN(c *Chip8) {
	var nByteJump uint16 = 2
	if uint16(c.registers[(c.opcode&0x0F00)>>8]) == c.opcode&0x00FF {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// SNE Vx, byte
// Skip next instruction if Vx != kk.
//
// The interpreter compares register Vx to kk, and if they are not equal,
// increments the program counter by 2.
func opcode4XNN(c *Chip8) {
	var nByteJump uint16 = 2
	if uint16(c.registers[(c.opcode&0x0F00)>>8]) != c.opcode&0x00FF {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// SE Vx, Vy
// Skip next instruction if Vx = Vy.
//
// The interpreter compares register Vx to register Vy, and if they are equal,
// increments the program counter by 2.
func opcode5XY0(c *Chip8) {
	var nByteJump uint16 = 2
	if c.registers[(c.opcode&0x0F00)>>8] == c.registers[(c.opcode&0x00F0)>>4] {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// LD Vx, byte
// Set Vx = kk.
//
// The interpreter puts the value kk into register Vx.
func opcode6XNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = uint8(c.opcode & 0x00FF)
	c.pc += 2
}

// ADD Vx, byte
// Set Vx = Vx + kk.
//
// Adds the value kk to the value of register Vx, then stores the result in Vx.
func opcode7XNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = uint8(c.registers[(c.opcode&0x0F00)>>8]) + uint8(c.opcode&0x00FF)
	c.pc += 2
}

// LD Vx, Vy
// Set Vx = Vy.
//
// Stores the value of register Vy in register Vx.
func opcode8XY0(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// OR Vx, Vy
// Set Vx = Vx OR Vy.
//
// Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx.
// A bitwise OR compares the corrseponding bits from two values, and if either bit is 1,
// then the same bit in the result is also 1. Otherwise, it is 0.
func opcode8XY1(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] | c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// AND Vx, Vy
// Set Vx = Vx AND Vy.
//
// Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
// A bitwise AND compares the corrseponding bits from two values, and if both bits are 1,
// then the same bit in the result is also 1. Otherwise, it is 0.
func opcode8XY2(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] & c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// XOR Vx, Vy
// Set Vx = Vx XOR Vy.
//
// Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx.
// An exclusive OR compares the corrseponding bits from two values,
// and if the bits are not both the same, then the corresponding bit in the result is set to 1.
// Otherwise, it is 0.
func opcode8XY3(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] ^ c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// ADD Vx, Vy
// Set Vx = Vx + Vy, set VF = carry.
//
// The values of Vx and Vy are added together.
// If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0.
// Only the lowest 8 bits of the result are kept, and stored in Vx.
func opcode8XY4(c *Chip8) {
	var carry uint8 = 0
	if c.registers[(c.opcode&0x00F0)>>4] > 0xFF-c.registers[(c.opcode&0x0F00)>>8] {
		carry = 1
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] + c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// SUB Vx, Vy
// Set Vx = Vx - Vy, set VF = NOT borrow.
//
// If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx,
// and the results stored in Vx.
func opcode8XY5(c *Chip8) {
	var carry uint8 = 1
	if c.registers[(c.opcode&0x00F0)>>4] > c.registers[(c.opcode&0x0F00)>>8] {
		carry = 0
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] - c.registers[(c.opcode&0x00F0)>>4]
	c.pc += 2
}

// SHR Vx {, Vy}
// Set Vx = Vx SHR 1.
//
// If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0.
// Then Vx is divided by 2.
func opcode8XY6(c *Chip8) {
	x := (c.opcode & 0x0F00) >> 8
	c.registers[0xF] = c.registers[x] & 0x1
	c.registers[x] >>= 1
	c.pc += 2
}

// SUBN Vx, Vy
// Set Vx = Vy - Vx, set VF = NOT borrow.
//
// If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy,
// and the results stored in Vx.
func opcode8XY7(c *Chip8) {
	var carry uint8 = 1
	if c.registers[(c.opcode&0x0F00)>>8] > c.registers[(c.opcode&0x00F0)>>4] {
		carry = 0
	}
	c.registers[0xF] = carry
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x00F0)>>4] - c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

// SHL Vx {, Vy}
// Set Vx = Vx SHL 1.
//
// If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0.
// Then Vx is multiplied by 2.
func opcode8XYE(c *Chip8) {
	c.registers[0xF] = c.registers[(c.opcode&0x0F00)>>8] >> 7
	c.registers[(c.opcode&0x0F00)>>8] = c.registers[(c.opcode&0x0F00)>>8] << 1
	c.pc += 2
}

// SNE Vx, Vy
// Skip next instruction if Vx != Vy.
//
// The values of Vx and Vy are compared, and if they are not equal,
// the program counter is increased by 2.
func opcode9XY0(c *Chip8) {
	var nByteJump uint16 = 2
	if c.registers[(c.opcode&0x0F00)>>8] != c.registers[(c.opcode&0x00F0)>>4] {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// LD I, addr
// Set I = nnn.
//
// The value of register I is set to nnn.
func opcodeANNN(c *Chip8) {
	c.i = c.opcode & 0x0FFF
	c.pc += 2
}

// JP V0, addr
// Jump to location nnn + V0.
//
// The program counter is set to nnn plus the value of V0.
func opcodeBNNN(c *Chip8) {
	c.pc = (c.opcode & 0x0FFF) + uint16(c.registers[0x0])
}

// RND Vx, byte
// Set Vx = random byte AND kk.
//
// The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk.
// The results are stored in Vx. See instruction 8xy2 for more information on AND.
func opcodeCXNN(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = uint8(rand.Intn(256)) & uint8(c.opcode&0x00FF)
	c.pc = c.pc + 2
}

// DRW Vx, Vy, nibble
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
//
// The interpreter reads n bytes from memory, starting at the address stored in I.
// These bytes are then displayed as sprites on screen at coordinates (Vx, Vy).
// Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1,
// otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates
// of the display, it wraps around to the opposite side of the screen.
// See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information
// on the Chip-8 screen and sprites.
func opcodeDXYN(c *Chip8) {
	x := uint16(c.registers[(c.opcode&0x0F00)>>8])
	y := uint16(c.registers[(c.opcode&0x00F0)>>4])
	height := c.opcode & 0x000F
	c.registers[0xF] = 0
	var yLine uint16
	var xLine uint16
	for yLine = 0; yLine < height; yLine++ {
		pixel := c.memory[c.i+yLine]
		for xLine = 0; xLine < 8; xLine++ {
			if (pixel & (0x80 >> xLine)) != 0 {
				// Avoid overflowing chip8's memory
				if x+xLine+((y+yLine)*64) >= gfxSize {
					continue
				}
				if c.gfx[(x+xLine+((y+yLine)*64))] == 1 {
					c.registers[0xF] = 1
				}
				c.gfx[x+xLine+((y+yLine)*64)] ^= 1
			}
		}
	}
	c.draw = true
	c.pc += 2
}

// SKP Vx
// Skip next instruction if key with the value of Vx is pressed.
//
// Checks the keyboard, and if the key corresponding to the value of Vx is currently
// in the down position, PC is increased by 2.
func opcodeEX9E(c *Chip8) {
	var nByteJump uint16 = 2
	if c.key[c.registers[(c.opcode&0x0F00)>>8]] == 1 {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// SKNP Vx
// Skip next instruction if key with the value of Vx is not pressed.
//
// Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position,
// PC is increased by 2.
func opcodeEXA1(c *Chip8) {
	var nByteJump uint16 = 2
	if c.key[c.registers[(c.opcode&0x0F00)>>8]] == 0 {
		nByteJump += 2
	}
	c.pc += nByteJump
}

// LD Vx, DT
// Set Vx = delay timer value.
//
// The value of DT is placed into Vx.
func opcodeFX07(c *Chip8) {
	c.registers[(c.opcode&0x0F00)>>8] = c.delayTimer
	c.pc += 2
}

// LD Vx, K
// Wait for a key press, store the value of the key in Vx.
//
// All execution stops until a key is pressed, then the value of that key is stored in Vx.
func opcodeFX0A(c *Chip8) {
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

// LD DT, Vx
// Set delay timer = Vx.
//
// DT is set equal to the value of Vx.
func opcodeFX15(c *Chip8) {
	c.delayTimer = c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

// LD ST, Vx
// Set sound timer = Vx.
//
// ST is set equal to the value of Vx.
func opcodeFX18(c *Chip8) {
	c.soundTimer = c.registers[(c.opcode&0x0F00)>>8]
	c.pc += 2
}

// ADD I, Vx
// Set I = I + Vx.
//
// The values of I and Vx are added, and the results are stored in I.
func opcodeFX1E(c *Chip8) {
	var carry uint8 = 0
	x, overflown := utils.AddU16(c.i, uint16(c.registers[(c.opcode&0x0F00)>>8]))
	if overflown || x > 0xFFF {
		carry = 1
	}
	c.registers[0xF] = carry
	c.i = x
	c.pc += 2
}

// LD F, Vx
// Set I = location of sprite for digit Vx.
//
// The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx.
// See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
func opcodeFX29(c *Chip8) {
	c.i = uint16(c.registers[(c.opcode&0x0F00)>>8]) * 0x5
	c.pc += 2
}

// LD B, Vx
// Store BCD representation of Vx in memory locations I, I+1, and I+2.
//
// The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location
// in I, the tens digit at location I+1, and the ones digit at location I+2.
func opcodeFX33(c *Chip8) {
	b := uint16(0)

	// perform 8 shifts
	for i := uint(0); i < 8; i++ {
		if (b>>0)&0xF >= 5 {
			b += 3
		}
		if (b>>4)&0xF >= 5 {
			b += 3 << 4
		}
		if (b>>8)&0xF >= 5 {
			b += 3 << 8
		}

		// apply shift, pull next bit
		b = (b << 1) | uint16(c.registers[(c.opcode&0x0F00)>>8]>>(7-i)&1)
	}

	// write to memory
	c.memory[c.i+0] = (b >> 8) & 0xF
	c.memory[c.i+1] = (b >> 4) & 0xF
	c.memory[c.i+2] = (b >> 0) & 0xF
	c.pc += 2
}

// LD [I], Vx
// Store registers V0 through Vx in memory starting at location I.
//
// The interpreter copies the values of registers V0 through Vx into memory, starting at the address in I.
func opcodeFX55(c *Chip8) {
	for i := 0; i < int((c.opcode&0x0F00)>>8)+1; i++ {
		c.memory[uint16(i)+c.i] = uint16(c.registers[i])
	}
	c.pc += 2
}

// LD Vx, [I]
// Read registers V0 through Vx from memory starting at location I.
//
// The interpreter reads values from memory starting at location I into registers V0 through Vx.
func opcodeFX65(c *Chip8) {
	for i := 0; i < int((c.opcode&0x0F00)>>8)+1; i++ {
		c.registers[i] = uint8(c.memory[c.i+uint16(i)])
	}
	c.pc += 2
}

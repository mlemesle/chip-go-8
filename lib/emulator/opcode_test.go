package emulator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func initChip8() *Chip8 {
	c := &Chip8{}
	c.Initialize()
	return c
}

func TestOpcode_00EE(t *testing.T) {
	c := initChip8()
	c.sp = 1
	c.stack[0] = 40
	opcode_00EE(c)
	assert.Equal(t, uint8(0), c.sp)
	assert.Equal(t, uint16(42), c.pc)
}

func TestOpcode_00E0(t *testing.T) {
	c := initChip8()
	opcode_00E0(c)
	assert.Equal(t, uint16(0x202), c.pc)
	for _, p := range c.gfx {
		assert.Equal(t, uint8(0), p)
	}
	assert.Equal(t, true, c.draw)
}

func TestOpcode_0NNN(t *testing.T) {
	c := initChip8()
	opcode_0NNN(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_1NNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0x1B0B
	opcode_1NNN(c)
	assert.Equal(t, uint16(0xB0B), c.pc)
}

func TestOpcode_2NNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0x2B0B
	opcode_2NNN(c)
	assert.Equal(t,
		[stackSize]uint16{0x200, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		c.stack)
	assert.Equal(t, uint8(1), c.sp)
	assert.Equal(t, uint16(0xB0B), c.pc)
}

func TestOpcode_3XNN_false(t *testing.T) {
	c := initChip8()
	c.opcode = 0x3ABB
	c.registers[0x0A] = 0xAA
	opcode_3XNN(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_3XNN_true(t *testing.T) {
	c := initChip8()
	c.opcode = 0x3ABB
	c.registers[0x0A] = 0xBB
	opcode_3XNN(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_4XNN_false(t *testing.T) {
	c := initChip8()
	c.opcode = 0x3ABB
	c.registers[0x0A] = 0xAA
	opcode_4XNN(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_4XNN_true(t *testing.T) {
	c := initChip8()
	c.opcode = 0x3ABB
	c.registers[0x0A] = 0xBB
	opcode_4XNN(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_5XY0_false(t *testing.T) {
	c := initChip8()
	c.opcode = 0x5AB0
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0xBB
	opcode_5XY0(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_5XY0_true(t *testing.T) {
	c := initChip8()
	c.opcode = 0x5AB0
	c.registers[0xA] = 0xBB
	c.registers[0xB] = 0xBB
	opcode_5XY0(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_6XNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0x6ABB
	opcode_6XNN(c)
	assert.Equal(t, uint8(0xBB), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_7XNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0x7ABB
	c.registers[0xA] = uint8(0x11)
	opcode_7XNN(c)
	assert.Equal(t, uint8(0xCC), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_7XNN_overflow(t *testing.T) {
	c := initChip8()
	c.opcode = 0x7AEE
	c.registers[0xA] = uint8(0x12)
	opcode_7XNN(c)
	assert.Equal(t, uint8(0x00), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY0(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB0
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0xBB
	opcode_8XY0(c)
	assert.Equal(t, uint8(0xBB), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY1(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB1
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0xBB
	opcode_8XY1(c)
	assert.Equal(t, uint8(0xAA|0xBB), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY2(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB2
	c.registers[0xA] = 0xEE
	c.registers[0xB] = 0x55
	opcode_8XY2(c)
	assert.Equal(t, uint8(0xEE&0x55), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY3(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB3
	c.registers[0xA] = 0xEE
	c.registers[0xB] = 0x55
	opcode_8XY3(c)
	assert.Equal(t, uint8(0xEE^0x55), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY4(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB4
	c.registers[0xA] = 0x11
	c.registers[0xB] = 0xAA
	opcode_8XY4(c)
	assert.Equal(t, uint8(0xBB), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY4_withCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB4
	c.registers[0xA] = 0x12
	c.registers[0xB] = 0xFF
	opcode_8XY4(c)
	assert.Equal(t, uint8(0x11), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY4_withAndWithoutCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB4
	c.registers[0xA] = 0x12
	c.registers[0xB] = 0xFF
	opcode_8XY4(c)
	assert.Equal(t, uint8(0x11), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)

	c.registers[0xB] = 0x22
	opcode_8XY4(c)
	assert.Equal(t, uint8(0x33), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_8XY5(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB5
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0x11
	opcode_8XY5(c)
	assert.Equal(t, uint8(0x99), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY5_withCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB5
	c.registers[0xA] = 0x10
	c.registers[0xB] = 0x22
	opcode_8XY5(c)
	assert.Equal(t, uint8(0xEE), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY5_withAndWithoutCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB5
	c.registers[0xA] = 0x10
	c.registers[0xB] = 0x22
	opcode_8XY5(c)
	assert.Equal(t, uint8(0xEE), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)

	c.registers[0xB] = 0x22
	opcode_8XY5(c)
	assert.Equal(t, uint8(0xCC), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_8XY6(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB6
	c.registers[0xA] = 0xFF
	opcode_8XY6(c)
	assert.Equal(t, uint8(0xFF>>1), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY7(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB7
	c.registers[0xA] = 0x11
	c.registers[0xB] = 0xAA
	opcode_8XY7(c)
	assert.Equal(t, uint8(0x99), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY7_withCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB7
	c.registers[0xA] = 0x22
	c.registers[0xB] = 0x10
	opcode_8XY7(c)
	assert.Equal(t, uint8(0xEE), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_8XY7_withAndWithoutCarry(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8AB7
	c.registers[0xA] = 0x22
	c.registers[0xB] = 0x10
	opcode_8XY7(c)
	assert.Equal(t, uint8(0xEE), c.registers[0xA])
	assert.Equal(t, uint8(0x00), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)

	c.registers[0xB] = 0xFF
	opcode_8XY7(c)
	assert.Equal(t, uint8(0x11), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_8XYE(t *testing.T) {
	c := initChip8()
	c.opcode = 0x8ABE
	c.registers[0xA] = 0xEF
	opcode_8XYE(c)
	assert.Equal(t, uint8(0xDE), c.registers[0xA])
	assert.Equal(t, uint8(0x01), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_9XY0_notEqual(t *testing.T) {
	c := initChip8()
	c.opcode = 0x9AB0
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0xBB
	opcode_9XY0(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_9XY0_equal(t *testing.T) {
	c := initChip8()
	c.opcode = 0x9AB0
	c.registers[0xA] = 0xAA
	c.registers[0xB] = 0xAA
	opcode_9XY0(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_ANNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0xA777
	opcode_ANNN(c)
	assert.Equal(t, uint16(0x777), c.i)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_BNNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0xB777
	c.registers[0x0] = 0x11
	opcode_BNNN(c)
	assert.Equal(t, uint16(0x788), c.pc)
}

func TestOpcode_CXNN(t *testing.T) {
	c := initChip8()
	c.opcode = 0xC7F0
	opcode_CXNN(c)
	assert.Equal(t, uint8(0x0), c.registers[0x7]&0x0F)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_DXYN(t *testing.T) {
	c := initChip8()
	c.opcode = 0xDABC
	opcode_DXYN(c)
	assert.Equal(t, uint16(0x202), c.pc)
	assert.Equal(t, true, c.draw)
}

func TestOpcode_EX9E_pressed(t *testing.T) {
	c := initChip8()
	c.opcode = 0xEA9E
	c.registers[0xA] = 0x07
	c.key[0x7] = 1
	opcode_EX9E(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_EX9E_not_pressed(t *testing.T) {
	c := initChip8()
	c.opcode = 0xEA9E
	c.registers[0xA] = 0x07
	c.key[0x7] = 0
	opcode_EX9E(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_EXA1_pressed(t *testing.T) {
	c := initChip8()
	c.opcode = 0xEA9E
	c.registers[0xA] = 0x07
	c.key[0x7] = 1
	opcode_EXA1(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_EXA1_not_pressed(t *testing.T) {
	c := initChip8()
	c.opcode = 0xEA9E
	c.registers[0xA] = 0x07
	c.key[0x7] = 0
	opcode_EXA1(c)
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_FX07(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA07
	c.registers[0xA] = 0x07
	c.delayTimer = 0x12
	opcode_FX07(c)
	assert.Equal(t, uint8(0x12), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX0A(t *testing.T) {
	c := initChip8()
	c.opcode = 0xEA9E
	c.registers[0xA] = 0x07
	opcode_FX0A(c)
	assert.Equal(t, uint8(0x07), c.registers[0xA])
	assert.Equal(t, uint16(0x200), c.pc)

	c.key[0x2] = 1
	opcode_FX0A(c)
	assert.Equal(t, uint8(0x02), c.registers[0xA])
	assert.Equal(t, uint16(0x202), c.pc)

	c.key[0x2] = 0
	c.key[0x8] = 1
	opcode_FX0A(c)
	assert.Equal(t, uint8(0x08), c.registers[0xA])
	assert.Equal(t, uint16(0x204), c.pc)
}

func TestOpcode_FX15(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA15
	c.registers[0xA] = 0x77
	opcode_FX15(c)
	assert.Equal(t, uint8(0x77), c.delayTimer)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX18(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA18
	c.registers[0xA] = 0x77
	opcode_FX18(c)
	assert.Equal(t, uint8(0x77), c.soundTimer)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX1E_without_carry(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA1E
	c.registers[0xA] = 0x11
	c.i = 0xAA
	opcode_FX1E(c)
	assert.Equal(t, uint16(0xBB), c.i)
	assert.Equal(t, uint8(0x0), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX1E_with_carry(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA1E
	c.registers[0xA] = 0x23
	c.i = 0xFFEE
	opcode_FX1E(c)
	assert.Equal(t, uint16(0x11), c.i)
	assert.Equal(t, uint8(0x1), c.registers[0xF])
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX29(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA29
	c.registers[0xA] = 0x11
	c.i = 0xAA
	opcode_FX29(c)
	assert.Equal(t, uint16(0x55), c.i)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX33(t *testing.T) {
	c := initChip8()
	c.opcode = 0xFA33
	opcode_FX33(c)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX55(t *testing.T) {
	c := initChip8()
	c.opcode = 0xF333
	c.registers[0] = 0x00
	c.registers[1] = 0x11
	c.registers[2] = 0x22
	c.registers[3] = 0x33
	c.i = 0xAA
	c.memory[0xAA+0x4] = 0xFF

	opcode_FX55(c)

	assert.Equal(t, uint16(0x00), c.memory[0xAA+0])
	assert.Equal(t, uint16(0x11), c.memory[0xAA+1])
	assert.Equal(t, uint16(0x22), c.memory[0xAA+2])
	assert.Equal(t, uint16(0x33), c.memory[0xAA+3])
	assert.Equal(t, uint16(0xFF), c.memory[0xAA+4])
	assert.Equal(t, uint16(0xAA), c.i)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestOpcode_FX65(t *testing.T) {
	c := initChip8()
	c.opcode = 0xF333
	c.i = 0xAA
	c.memory[c.i+0] = 0x00
	c.memory[c.i+1] = 0x11
	c.memory[c.i+2] = 0x22
	c.memory[c.i+3] = 0x33
	c.registers[0x4] = 0xFF

	opcode_FX65(c)

	assert.Equal(t, uint8(0x00), c.registers[0])
	assert.Equal(t, uint8(0x11), c.registers[1])
	assert.Equal(t, uint8(0x22), c.registers[2])
	assert.Equal(t, uint8(0x33), c.registers[3])
	assert.Equal(t, uint8(0xFF), c.registers[4])
	assert.Equal(t, uint16(0xAA), c.i)
	assert.Equal(t, uint16(0x202), c.pc)
}

func TestReturnAfterCall(t *testing.T) {
	c := initChip8()
	c.opcode = 0x2B0B
	opcode_2NNN(c)
	opcode_00EE(c)
	assert.Equal(t,
		[stackSize]uint16{0x200, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		c.stack)
	assert.Equal(t, byte(0), c.sp)
	assert.Equal(t, uint16(0x202), c.pc)
}

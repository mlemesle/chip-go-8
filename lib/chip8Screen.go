package lib

import (
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	// "strings"
)

type Chip8Screen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	w        int32
	h        int32
	ratio    int32
}

type Chip8Scren interface {
	Init(c *Chip8, w, h, ratio int32) error
	Destroy()
	Draw(gfx [gfxSize]uint16) error
	HandleEvent(c *Chip8) bool
}

func (c8s *Chip8Screen) Init(c *Chip8, w, h, ratio int32) error {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("Chip 8 emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 62*ratio, 32*ratio, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)
	if err != nil {
		return err
	}
	c8s.window = window
	c8s.renderer = renderer
	c8s.w = w
	c8s.h = h
	c8s.ratio = ratio
	return nil
}

func (c8s *Chip8Screen) Destroy() {
	c8s.renderer.Destroy()
	c8s.window.Destroy()
}

func (c8s *Chip8Screen) Draw(gfx [gfxSize]uint8) error {
	c8s.renderer.SetDrawColor(0, 0, 0, 255)
	if err := c8s.renderer.Clear(); err != nil {
		return err
	}

	var x, y int32
	for y = 0; y < c8s.h; y++ {
		for x = 0; x < c8s.w; x++ {
			// Values of pixel are stored in 1D array of size 64 * 32
			if gfx[x+y*64] == 0 {
				continue
			}
			c8s.renderer.SetDrawColor(255, 255, 255, 255)
			if err := c8s.renderer.FillRect(&sdl.Rect{
				X: int32(x) * c8s.ratio,
				Y: int32(y) * c8s.ratio,
				W: c8s.ratio,
				H: c8s.ratio,
			}); err != nil {
				return err
			}
		}
	}

	c8s.renderer.Present()
	return nil
}

func (c8s *Chip8Screen) HandleEvent(c *Chip8) bool {
	// Poll for Quit and Keyboard events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch et := event.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.KeyboardEvent:
			if et.Type == sdl.KEYUP {
				switch et.Keysym.Sym {
				case sdl.K_1:
					c.SetKeyUp(0x1)
				case sdl.K_2:
					c.SetKeyUp(0x2)
				case sdl.K_3:
					c.SetKeyUp(0x3)
				case sdl.K_4:
					c.SetKeyUp(0xC)
				case sdl.K_a:
					c.SetKeyUp(0x4)
				case sdl.K_z:
					c.SetKeyUp(0x5)
				case sdl.K_e:
					c.SetKeyUp(0x6)
				case sdl.K_r:
					c.SetKeyUp(0xD)
				case sdl.K_q:
					c.SetKeyUp(0x7)
				case sdl.K_s:
					c.SetKeyUp(0x8)
				case sdl.K_d:
					c.SetKeyUp(0x9)
				case sdl.K_f:
					c.SetKeyUp(0xE)
				case sdl.K_w:
					c.SetKeyUp(0xA)
				case sdl.K_x:
					c.SetKeyUp(0x0)
				case sdl.K_c:
					c.SetKeyUp(0xB)
				case sdl.K_v:
					c.SetKeyUp(0xF)
				}
			} else if et.Type == sdl.KEYDOWN {
				switch et.Keysym.Sym {
				case sdl.K_1:
					c.SetKeyDown(0x1)
				case sdl.K_2:
					c.SetKeyDown(0x2)
				case sdl.K_3:
					c.SetKeyDown(0x3)
				case sdl.K_4:
					c.SetKeyDown(0xC)
				case sdl.K_a:
					c.SetKeyDown(0x4)
				case sdl.K_z:
					c.SetKeyDown(0x5)
				case sdl.K_e:
					c.SetKeyDown(0x6)
				case sdl.K_r:
					c.SetKeyDown(0xD)
				case sdl.K_q:
					c.SetKeyDown(0x7)
				case sdl.K_s:
					c.SetKeyDown(0x8)
				case sdl.K_d:
					c.SetKeyDown(0x9)
				case sdl.K_f:
					c.SetKeyDown(0xE)
				case sdl.K_w:
					c.SetKeyDown(0xA)
				case sdl.K_x:
					c.SetKeyDown(0x0)
				case sdl.K_c:
					c.SetKeyDown(0xB)
				case sdl.K_v:
					c.SetKeyDown(0xF)
				}
			}
		}
	}
	return false
}

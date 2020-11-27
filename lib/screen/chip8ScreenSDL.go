package screen

import (
	"github.com/mlemesle/chip-go-8/lib/emulator"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

// Chip8ScreenSDL represents a display for the chip8, it uses the SDL library
type Chip8ScreenSDL struct {
	durationBetweenFrame time.Duration
	lastDraw             time.Time
	window               *sdl.Window
	renderer             *sdl.Renderer
	w                    int32
	h                    int32
	ratio                int32
}

// NewChip8ScreenSDL creates a new non-initialized Chip8ScreenSDL
func NewChip8ScreenSDL(w, h, ratio int32) *Chip8ScreenSDL {
	return &Chip8ScreenSDL{
		durationBetweenFrame: time.Duration(time.Second / 60),
		w:                    w,
		h:                    h,
		ratio:                ratio,
	}
}

// Init initializes the given Chip8ScreenSDL
func (c8s *Chip8ScreenSDL) Init() error {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("Chip 8 emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, c8s.w*c8s.ratio, c8s.h*c8s.ratio, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)
	if err != nil {
		return err
	}
	c8s.window = window
	c8s.renderer = renderer
	c8s.lastDraw = time.Now()
	return nil
}

// Destroy cleans the struct and free the memory
func (c8s *Chip8ScreenSDL) Destroy() {
	c8s.renderer.Destroy()
	c8s.window.Destroy()
}

// Draw displays the gfx of the Chip8 on the screen
func (c8s *Chip8ScreenSDL) Draw(c *emulator.Chip8) error {
	if elapsed := time.Since(c8s.lastDraw); elapsed < c8s.durationBetweenFrame {
		timeToWait := time.Until(c8s.lastDraw.Add(c8s.durationBetweenFrame - elapsed))
		time.Sleep(timeToWait)
	}

	c8s.renderer.SetDrawColor(0, 0, 0, 255)
	if err := c8s.renderer.Clear(); err != nil {
		return err
	}

	var x, y int32
	for y = 0; y < c8s.h; y++ {
		for x = 0; x < c8s.w; x++ {
			if c.GetGFX()[x+y*64] == 0 {
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
	c.SetDraw(false)
	c8s.lastDraw = time.Now()
	return nil
}

// HandleEvent processes the user's inputs
func (c8s *Chip8ScreenSDL) HandleEvent(c *emulator.Chip8) bool {
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

package main

import (
	"github.com/mlemesle/chip-go-8/lib/emulator"
	"github.com/mlemesle/chip-go-8/lib/screen"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	chip8 := emulator.New()
	chip8.Initialize()
	err := chip8.LoadMemory("rom/pong.c8")
	if err != nil {
		panic(err)
	}

	chip8ScreenSDL := screen.NewChip8ScreenSDL(64, 32, 20)
	err = chip8ScreenSDL.Init()
	if err != nil {
		panic(err)
	}
	defer chip8ScreenSDL.Destroy()

	for {
		if err = chip8.EmulateCycle(); err != nil {
			panic(err)
		}

		if chip8.NeedDraw() {
			if err = chip8ScreenSDL.Draw(chip8); err != nil {
				panic(err)
			}
		}

		quitEvent := chip8ScreenSDL.HandleEvent(chip8)
		if quitEvent {
			os.Exit(0)
		}
		sdl.Delay(1000 / 1000)
	}
}

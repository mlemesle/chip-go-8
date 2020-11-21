package main

import (
	lib "github.com/mlemesle/chip-go-8/lib"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	chip8 := &lib.Chip8{}
	chip8.Initialize()
	err := chip8.LoadMemory("rom/pong.c8")
	if err != nil {
		panic(err)
	}

	chip8Screen := lib.Chip8Screen{}
	err = chip8Screen.Init(chip8, 64, 32, 20)
	if err != nil {
		panic(err)
	}
	defer chip8Screen.Destroy()

	for {
		if err = chip8.EmulateCycle(); err != nil {
			panic(err)
		}

		if chip8.NeedDraw() {
			if err = chip8Screen.Draw(chip8.GetGFX()); err != nil {
				panic(err)
			}
		}

		quitEvent := chip8Screen.HandleEvent(chip8)
		if quitEvent {
			os.Exit(0)
		}

		sdl.Delay(1000 / 60)
	}
}

package main

import (
	"flag"
	"github.com/mlemesle/chip-go-8/lib/beeper"
	"github.com/mlemesle/chip-go-8/lib/emulator"
	"github.com/mlemesle/chip-go-8/lib/screen"
	"os"
)

func main() {
	ratio := flag.Int("ratio", 20, "The ratio of the screen. The screen standard size is 64x32.")
	isMuted := flag.Bool("mute", false, "The emulator will be muted if set.")
	runTest := flag.Bool("test", false, "If set, the emulator will boot with the test chip8 image from https://github.com/corax89/chip8-test-rom")
	romFile := flag.String("rom", "rom/pong.c8", "Specify a rom file to run. If not set, a pong image will be loaded")
	flag.Parse()

	chip8ScreenSDL := screen.NewChip8ScreenSDL(64, 32, int32(*ratio))
	err := chip8ScreenSDL.Init()
	if err != nil {
		panic(err)
	}
	defer chip8ScreenSDL.Destroy()

	var chip8Beeper beeper.BeeperInterface
	if *isMuted {
		chip8Beeper = beeper.NewMute()
	} else {
		chip8Beeper = beeper.NewSDL()
	}
	err = chip8Beeper.Init()
	if err != nil {
		panic(err)
	}
	defer chip8Beeper.Destroy()

	chip8 := emulator.New()
	chip8.Initialize(chip8Beeper)
	if *runTest {
		*romFile = "rom/test_opcode.ch8"
	}
	err = chip8.LoadMemory(*romFile)
	if err != nil {
		panic(err)
	}

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
	}
}

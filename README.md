# chip-go-8

![Pong game](https://github.com/mlemesle/chip-go-8/blob/master/pong.png "Pong game")

## Goal

The goal of this project is to improve my Go skills and to enter into the world of emulation.
So there it is, my Chip8 emulator !

## Prerequisite

To run this project you need :
* Go, at least v1.14 (you can install it from [here](https://golang.org/doc/install "Golang install page"))
* The SDL library (you can install it from [here](https://wiki.libsdl.org/Installation "SDL install page"))

If you're all set, you can now install this project !

## Installation

To retrieve this project, use

```
go get -u github.com/mlemesle/chip-go-8
```

Now go to the project's root directory (where main.go is located) and run `go build` to build the project and generate an executable. You now have chip-go-8 executable file in your current directory !

## Running the emulator

To launch the emulator, simply run 

```
./chip-go-8
```

It will run the emulator with default parameters :
* pong game will be loaded
* sound is active
* screen size is 1280x620px

You can customize the emulator using the following flags :

```
$ ./chip-go-8 --help
Usage of ./chip-go-8:
  -mute
    	The emulator will be muted if set.
  -ratio int
    	The ratio of the screen. The screen standard size is 64x32. (default 20)
  -rom string
    	Specify a rom file to run. If not set, a pong image will be loaded (default "rom/pong.c8")
  -test
    	If set, the emulator will boot with the test chip8 image from https://github.com/corax89/chip8-test-rom
```

So for example `./chip-go-8 -ratio 15 -rom path/to/file.c8 -mute` will run the emulator with a screen size of 960x480px, load the file located at path/to/file.c8 and won't produce any sound.

Feel free to try `./chip-go-8 -test`, it will run a special test image to assert that all opcodes are correctly implemented !

## Keyboard controls

> The computers which originally used the Chip-8 Language had a 16-key hexadecimal keypad with the following layout: *[original content](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#keyboard)*

```
1 | 2 | 3 | C 
4 | 5 | 6 | D 
7 | 8 | 9 | E 
A | 0 | B | F 
```

I remapped this keypad to 

```
1 | 2 | 3 | 4
A | Z | E | R 
Q | S | D | F 
W | X | C | V 
```

## Where to find roms

You can find pretty cool roms right [here](https://github.com/dmatlack/chip8) ! You just need to download one of them, pass it to chip-go-8 and you're ready to go !

## References

* [Wikipedia's Chip8 page](https://en.wikipedia.org/wiki/CHIP-8)
* [Chip8 technical reference v1.0](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
* [How to write a Chip8 emulator](http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/)
* [SDL library documentation](https://wiki.libsdl.org/APIByCategory)
* [Github page of the opcode test image](https://github.com/corax89/chip8-test-rom)
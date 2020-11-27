package beeper

// typedef unsigned char Uint8;
// void AudioCallback(void *userdata, Uint8 *stream, int len);
import "C"

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"reflect"
	"unsafe"
)

// SDL is a beeper using SDL library
type SDL struct {
	audioDeviceID sdl.AudioDeviceID
}

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	buf := *(*[]C.Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  n,
		Cap:  n,
	}))

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += 2 * math.Pi * 440 / 44100
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}

// NewSDL creates a new instance of a beeper using SDL library
func NewSDL() *SDL {
	return &SDL{}
}

// Init initializes the given SDL
func (b *SDL) Init() error {
	spec := &sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_S16SYS,
		Channels: 1,
		Samples:  2048,
		Callback: sdl.AudioCallback(C.AudioCallback),
	}
	audioDeviceID, err := sdl.OpenAudioDevice(sdl.GetAudioDeviceName(0, false), false, spec, &sdl.AudioSpec{}, 0)
	if err != nil {
		return err
	}

	b.audioDeviceID = audioDeviceID

	return nil
}

// Beep does a short beep
func (b *SDL) Beep() {
	sdl.PauseAudioDevice(b.audioDeviceID, false)
	sdl.Delay(50)
	sdl.PauseAudioDevice(b.audioDeviceID, true)
}

// Destroy free SDL allocated components
func (b *SDL) Destroy() {
	sdl.CloseAudioDevice(b.audioDeviceID)
}

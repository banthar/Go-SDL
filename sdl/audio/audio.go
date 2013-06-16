/*
 * Copyright neagix - Feb 2013
 * Copyright: ⚛ <0xe2.0x9a.0x9b@gmail.com> 2010
 * 
 *
 * The contents of this file can be used freely,
 * except for usages in immoral contexts.
 * 
 */

/*
An interface to low-level SDL sound functions
with support for callbacks and rudimentary stream mixing
*/
package audio

// #cgo pkg-config: sdl
// #cgo freebsd LDFLAGS: -lrt
// #cgo linux LDFLAGS: -lrt
// #cgo windows LDFLAGS: -lpthread
// #include <SDL_audio.h>
// #include "callback.h"
import "C"
import "unsafe"
import "reflect"

// The version of Go-SDL audio bindings.
// The version descriptor changes into a new unique string
// after a semantically incompatible Go-SDL update.
//
// The returned value can be checked by users of this package
// to make sure they are using a version with the expected semantics.
//
// If Go adds some kind of support for package versioning, this function will go away.
func GoSdlAudioVersion() string {
	return "Go-SDL audio 1.2"
}

var userDefinedCallback func(unsafe.Pointer, int)

// this is the callback called from the C code
// without any special glue since the CGO threads insanity has been fixed =)
//export streamCallback
func streamCallback(arg unsafe.Pointer) {
	ctx := (*C.context)(arg)

	// call the actual Go callback defined by user
	//NOTE: here buffer truncation possible with large NumBytes
	userDefinedCallback(ctx.Stream, int(ctx.NumBytes))
}

// Audio format
const (
	AUDIO_U8     = C.AUDIO_U8
	AUDIO_S8     = C.AUDIO_S8
	AUDIO_U16LSB = C.AUDIO_U16LSB
	AUDIO_S16LSB = C.AUDIO_S16LSB
	AUDIO_U16MSB = C.AUDIO_U16MSB
	AUDIO_S16MSB = C.AUDIO_S16MSB
	AUDIO_U16    = C.AUDIO_U16
	AUDIO_S16    = C.AUDIO_S16
)

// Native audio byte ordering
const (
	AUDIO_U16SYS = C.AUDIO_U16SYS
	AUDIO_S16SYS = C.AUDIO_S16SYS
)

type AudioSpec struct {
	Freq                int
	Format              uint16 // If in doubt, use AUDIO_S16SYS
	Channels            uint8  // 1 or 2
	Out_Silence         uint8
	Samples             uint16 // A power of 2, preferrably 2^11 (2048) or more
	Out_Size            uint32
	UserDefinedCallback func(unsafe.Pointer, int)
}

var alreadyOpened bool

const (
	// play concatenating samples
	AE_PLAY_CONCAT = iota
	//TODO: downmix mode
	AE_PAUSE
	AE_UNPAUSE
)

type AudioEvent struct {
	Event  int
	Buffer []int16

	AudioType int
}

// Audio status
const (
	SDL_AUDIO_STOPPED = C.SDL_AUDIO_STOPPED
	SDL_AUDIO_PLAYING = C.SDL_AUDIO_PLAYING
	SDL_AUDIO_PAUSED  = C.SDL_AUDIO_PAUSED
)

var PlayLoop chan AudioEvent
var PlayQueueSize int
var TailBufferS16 []int16

//NOTE: we assume this is not going to be called concurrently :)
func DownstreamPlaybackS16(buffer unsafe.Pointer, bufferSize int) {

	// hack to convert C void* pointer to proper Go slice
	var stream []int16
	streamLen := bufferSize / 2
	tailBufferLen := len(TailBufferS16)
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&stream)))
	sliceHeader.Cap = streamLen
	sliceHeader.Len = streamLen
	sliceHeader.Data = uintptr(buffer)

	// in theory we should overmix samples, but for that a double-buffered approach is best
	// right now we will just concatenate samples, quirky but serves the purpose of this game

	// first thing, pick any eventual remnants buffer to copy over
	dstOffset := 0
	if nil != TailBufferS16 && tailBufferLen > 0 {
		for i := range TailBufferS16 {
			stream[dstOffset] = TailBufferS16[i]
			dstOffset++

			// we have copied enough
			if dstOffset == streamLen {
				newTailBufferS16 := make([]int16, tailBufferLen-dstOffset)
				copy(newTailBufferS16, TailBufferS16[dstOffset:])
				TailBufferS16 = newTailBufferS16
				return
			}
		}

		// we have consumed the whole TailBufferS16, reset it
		TailBufferS16 = nil
	}

	// pick next beep object
	for {
		// if there's nothing in queue, just play what we got so far
		// by default stream is a silence buffer
		if 0 == PlayQueueSize {
			return
		}

		ae := <-PlayLoop

		switch ae.Event {
		case AE_UNPAUSE:
			{
				PauseAudio(false)
				continue
			}
		case AE_PAUSE:
			{
				PauseAudio(true)
				continue
			}
		}

		// prepare eventual tail buffer
		toBeCopied := streamLen - dstOffset
		overflowingSamples := len(ae.Buffer) - toBeCopied
		if overflowingSamples > 0 {
			TailBufferS16 = make([]int16, overflowingSamples)
		}

		copy(stream[dstOffset:], ae.Buffer)
		if overflowingSamples > 0 {
			copy(TailBufferS16, ae.Buffer[toBeCopied:])
		}

		// we have "eaten" a sound object from the queue
		PlayQueueSize--

		// this sound object fully satisfied the buffer
		if overflowingSamples > 0 {
			return
		}

		// in case of perfect boundary match
		if streamLen == dstOffset+toBeCopied {
			return
		}

		// loop again, pick another sound
	}
}

func OpenAudio(desired, obtained_orNil *AudioSpec) int {
	if alreadyOpened {
		panic("more than 1 audio stream currently not supported")
	}

	// copy handle to user-defined callback function, if defined
	// it is perfectly supported to not specify the callback function
	// in that case you will use default SendAudio semantics
	// note that if you specify a callback and use SendAudio, a hangup will instead happen
	// when calling SendAudio	

	if nil != desired.UserDefinedCallback {
		userDefinedCallback = desired.UserDefinedCallback
	} else {
		// default playback (16-bit signed)
		userDefinedCallback = DownstreamPlaybackS16
		PlayLoop = make(chan AudioEvent)
	}

	var C_desired, C_obtained *C.SDL_AudioSpec

	C_desired = new(C.SDL_AudioSpec)
	C_desired.freq = C.int(desired.Freq)
	C_desired.format = C.Uint16(desired.Format)
	C_desired.channels = C.Uint8(desired.Channels)
	C_desired.samples = C.Uint16(desired.Samples)
	// there is an unique C callback acting as proxy to the different Go callbacks
	// see streamContext()
	C_desired.callback = C.callback_getCallback()

	if obtained_orNil != nil {
		if desired != obtained_orNil {
			C_obtained = new(C.SDL_AudioSpec)
		} else {
			C_obtained = C_desired
		}
	}

	status := C.SDL_OpenAudio(C_desired, C_obtained)

	if status == 0 {
		alreadyOpened = true
	}

	if obtained_orNil != nil {
		obtained := obtained_orNil

		obtained.Freq = int(C_obtained.freq)
		obtained.Format = uint16(C_obtained.format)
		obtained.Channels = uint8(C_obtained.channels)
		obtained.Samples = uint16(C_obtained.samples)
		obtained.Out_Silence = uint8(C_obtained.silence)
		obtained.Out_Size = uint32(C_obtained.size)
	}

	return int(status)
}

func CloseAudio() {
	if !alreadyOpened {
		panic("SDL audio not opened")
	}

	PauseAudio(true)

	C.SDL_CloseAudio()
}

func GetAudioStatus() int {
	return int(C.SDL_GetAudioStatus())
}

// Pause or unpause the audio.
func PauseAudio(pause_on bool) {
	if pause_on {
		C.SDL_PauseAudio(1)
	} else {
		C.SDL_PauseAudio(0)
	}
}

// pause or unpause the audio synchronously
func PauseAudioSync(pause_on bool) {
	if nil == PlayLoop {
		panic("Cannot use PauseAudioSync with custom callback")
	}
	if pause_on {
		PlayLoop <- AudioEvent{Event: AE_PAUSE}
	} else {
		PlayLoop <- AudioEvent{Event: AE_UNPAUSE}
	}
}

// Send samples to the audio device (AUDIO_S16SYS format).
// This function blocks until all the samples are consumed by the SDL audio thread.
func SendAudio_int16(data []int16) {
	PlayQueueSize++
	PlayLoop <- AudioEvent{Event: AE_PLAY_CONCAT, Buffer: data, AudioType: AUDIO_S16}
}

// Send samples to the audio device (AUDIO_U16SYS format).
// This function blocks until all the samples are consumed by the SDL audio thread.
func SendAudio_uint16(data []uint16) {
	panic("Only int16 samples currently supported")
}

// Send samples to the audio device (AUDIO_S8 format).
// This function blocks until all the samples are consumed by the SDL audio thread.
func SendAudio_int8(data []int8) {
	panic("Only int16 samples currently supported")
}

// Send samples to the audio device (AUDIO_U8 format).
// This function blocks until all the samples are consumed by the SDL audio thread.
func SendAudio_uint8(data []uint8) {
	panic("Only int16 samples currently supported")
}

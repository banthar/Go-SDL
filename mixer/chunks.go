/*
A binding of Mix_Chunk.

The binding works pretty much the same as the original, although a few
functions have been changed to be in a more object-oriented style
(eg. Rather than mixer.FreeChunk(sound) it's sound.Free() )
*/
package mixer

// #cgo pkg-config: sdl
// #cgo LDFLAGS: -lSDL_mixer
// #include "SDL_mixer.h"
import "C"
import "unsafe"

// A Chunk file.
type Chunk struct {
	cchunk *C.Mix_Chunk
}

// Loads a sound file to use.
func LoadWAV(file string) *Chunk {
	cfile := C.CString(file)
	rb := C.CString("rb")

	cchunk := C.Mix_LoadWAV_RW(C.SDL_RWFromFile(cfile, rb), 1)
	C.free(unsafe.Pointer(cfile))
	C.free(unsafe.Pointer(rb))

	if cchunk == nil {
		return nil
	}

	return &Chunk{cchunk}
}

// Frees the loaded sound file.
func (c *Chunk) Free() { C.Mix_FreeChunk(c.cchunk) }

//Returns: previous chunk volume setting, Sets volume of Chunk
func (c *Chunk) Volume(volume int) int {
	return int(C.Mix_VolumeChunk(c.cchunk, C.int(volume)))
}

//Play chunk on channel, or if channel is -1, pick the ﬁrst free unreserved channel. The sample
//will play for loops+1 number of times, unless stopped by halt, or fade out, or setting a new
//expiration time of less time than it would have originally taken to play the loops, or closing
//the mixer
//Returns: the channel the sample is played on. On any errors, -1 is returne
func (c *Chunk) PlayChannel(channel, loops int) int {
	return c.PlayChannelTimed(channel, loops, -1)
}

//If the sample is long enough and has enough loops then the sample will stop after ticks milliseconds.
//Otherwise this function is the same as chunk.PlayChannel()
//Returns: the channel the sample is played on. On any errors, -1 is returned
func (c *Chunk) PlayChannelTimed(channel, loops, ticks int) int {
	return int(C.Mix_PlayChannelTimed(C.int(channel), c.cchunk, C.int(loops), C.int(ticks)))
}

//Play chunk on channel, or if channel is -1, pick the ﬁrst free unreserved channel.
//The channel volume starts at 0 and fades up to full volume over ms milliseconds of time.
//The sample may end before the fade-in is complete if it is too short or doesn’t have enough
//loops. The sample will play for loops+1 number of times, unless stopped.
//Returns: the channel the sample is played on. On any errors, -1 is returned
func (c *Chunk) FadeInChannel(channel, loops, ms int) int {
	return c.FadeInChannelTimed(channel, loops, ms, -1)
}

//If the sample is long enough and has enough loops then the sample will stop after ticks
//milliseconds.
//Returns: the channel the sample is played on. On any errors, -1 is returned.
func (c *Chunk) FadeInChannelTimed(channel, loops, ms, ticks int) int {
	return int(C.Mix_FadeInChannelTimed(C.int(channel), c.cchunk, C.int(loops), C.int(ms), C.int(ticks)))
}

//Returns: Pointer to the Mix Chunk. nil is returned if the channel is not allocated, or
//if the channel has not played any samples yet
func GetChunk(channel int) *Chunk {
	out := C.Mix_GetChunk(C.int(channel))
	if out == nil {
		return nil
	}
	return &Chunk{out}
}

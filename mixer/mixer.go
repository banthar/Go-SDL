/*
A binding of SDL_mixer.

The binding works pretty much the same as the original, although a few
functions have been changed to be in a more object-oriented style
(eg. Rather than mixer.FreeMusic(song) it's song.Free() )
*/
package mixer

// #cgo pkg-config: SDL_mixer
// #include <SDL_mixer.h>
import "C"
import "unsafe"

// A music file.
type Music struct {
	cmusic *C.Mix_Music
}

// Initializes SDL_mixer.  Return 0 if successful and -1 if there were
// initialization errors.
func OpenAudio(frequency int, format uint16, channels, chunksize int) int {
	return int(C.Mix_OpenAudio(C.int(frequency), C.Uint16(format),
		C.int(channels), C.int(chunksize)))
}

// Queries the mixer format. Returns (0,0,0) if audio has not been
// opened, and (frequency, format, channels) otherwise
func QuerySpec() (int, uint16, int) {
	var frequency C.int
	var format C.Uint16
	var channels C.int
	if C.Mix_QuerySpec(&frequency, &format, &channels) == 0 {
		return 0, 0, 0
	}
	return int(frequency), uint16(format), int(channels)
}

// Shuts down SDL_mixer.
func CloseAudio() { C.Mix_CloseAudio() }

// Loads a music file to use.
func LoadMUS(file string) *Music {
	cfile := C.CString(file)
	cmusic := C.Mix_LoadMUS(cfile)
	C.free(unsafe.Pointer(cfile))

	if cmusic == nil {
		return nil
	}

	return &Music{cmusic}
}

// Frees the loaded music file.
func (m *Music) Free() { C.Mix_FreeMusic(m.cmusic) }

// Play the music and loop a specified number of times.  Passing -1 makes
// the music loop continuously.
func (m *Music) PlayMusic(loops int) int {
	return int(C.Mix_PlayMusic(m.cmusic, C.int(loops)))
}

// Play the music and loop a specified number of times.  During the first loop,
// fade in for the milliseconds specified.  Passing -1 makes the music loop
// continuously.  The fade-in effect only occurs during the first loop.
func (m *Music) FadeInMusic(loops, ms int) int {
	return int(C.Mix_FadeInMusic(m.cmusic, C.int(loops), C.int(ms)))
}

// Same as FadeInMusic, only with a specified position to start the music at.
func (m *Music) FadeInMusicPos(loops, ms int, position float64) int {
	return int(C.Mix_FadeInMusicPos(m.cmusic, C.int(loops), C.int(ms),
		C.double(position)))
}

// Sets the volume to the value specified.
func VolumeMusic(volume int) int { return int(C.Mix_VolumeMusic(C.int(volume))) }

// Pauses the music playback.
func PauseMusic() { C.Mix_PauseMusic() }

// Unpauses the music.
func ResumeMusic() { C.Mix_ResumeMusic() }

// Rewinds music to the start.
func RewindMusic() { C.Mix_RewindMusic() }

// Sets the position of the currently playing music.
func SetMusicPosition(position float64) int {
	return int(C.Mix_SetMusicPosition(C.double(position)))
}

// Halt playback of music.
func HaltMusic() { C.Mix_HaltMusic() }

// Fades out music over the milliseconds specified.  Music is halted after
// the fade out is completed.
func FadeOutMusic(ms int) int { return int(C.Mix_FadeOutMusic(C.int(ms))) }

// Returns the type of the currently playing music.
func GetMusicType() int { return int(C.Mix_GetMusicType(nil)) }

// Returns the type of the music.
func (m *Music) GetMusicType() int { return int(C.Mix_GetMusicType(m.cmusic)) }

// Returns 1 if music is currently playing and 0 if not.
func PlayingMusic() int { return int(C.Mix_PlayingMusic()) }

// Returns 1 if music is paused and 0 if not.
func PausedMusic() int { return int(C.Mix_PausedMusic()) }

// Tells you whether music is fading in, out, or not at all.
func FadingMusic() int { return int(C.Mix_FadingMusic()) }

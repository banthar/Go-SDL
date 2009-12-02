/*
A binding of SDL_mixer.
*/
package mixer

// #include <SDL/SDL_mixer.h>
import "C"
import "unsafe"

type Music struct {
	cmusic *C.Mix_Music;
}

// Initializes SDL_mixer.
func OpenAudio(frequency int, format uint16, channels int, chunksize int) int {
	return int(C.Mix_OpenAudio(C.int(frequency), C.Uint16(format),
		C.int(channels), C.int(chunksize)));
}

// Shuts down SDL_mixer.
func CloseAudio() {
	C.Mix_CloseAudio();
}

// Loads a music file to use.
func LoadMUS(file string) *Music {
	cfile := C.CString(file);
	cmusic := C.Mix_LoadMUS(cfile);
	C.free(unsafe.Pointer(cfile));
	return &Music{cmusic};
}

// Frees the loaded music file.
func (m *Music) Free() {
	C.Mix_FreeMusic(m.cmusic);
}

// Play the music and loop a specified number of times.  Passing -1 makes
// the music loop continuously.
func (m *Music) PlayMusic(loops int) int {
	return int(C.Mix_PlayMusic(m.cmusic, C.int(loops)));
}

// Play the music and loop a specified number of times.  During the first loop,
// fade in for the milliseconds specified.  Passing -1 makes the music loop
// continuously.  The fade-in effect only occurs during the first loop.
func (m *Music) FadeInMusic(loops int, ms int) int {
	return int(C.Mix_FadeInMusic(m.cmusic, C.int(loops), C.int(ms)));
}

// Same as FadeInMusic, only with a specified position to start the music at.
func (m *Music) FadeInMusicPos(loops int, ms int, position float) int {
	return int(C.Mix_FadeInMusicPos(m.cmusic, C.int(loops), C.int(ms),
		C.double(position)));
}

// Sets the volume to the value specified.
func VolumeMusic(volume int) int {
	return int(C.Mix_VolumeMusic(C.int(volume)));
}

// Pauses the music playback.
func PauseMusic() {
	C.Mix_PauseMusic();
}

// Unpauses the music.
func ResumeMusic() {
	C.Mix_ResumeMusic();
}

// Rewinds music to the start.
func RewindMusic() {
	C.Mix_RewindMusic();
}

// Sets the position of the currently playing music.
func SetMusicPosition(position float) int {
	return int(C.Mix_SetMusicPosition(C.double(position)));
}

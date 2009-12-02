/*
A binding of SDL_mixer.
*/
package mixer

// #include <SDL/SDL_mixer.h>
import "C"

type Music struct {
	cmusic = *C.Mix_Music;
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

func LoadMUS(file string) *Music {
	cfile := C.string(file);
	cmusic := C.Mix_LoadMUS(cfile);
	return &Music{cmusic};
}

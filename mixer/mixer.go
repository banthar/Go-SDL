/*
A binding of SDL_mixer.
*/
package mixer

// #include <SDL/SDL_mixer.h>
import "C"

// Initializes SDL_mixer.
func OpenAudio(frequency int, format uint16, channels int, chunksize int) int {
	return int(C.Mix_OpenAudio(C.int(frequency), C.Uint16(format),
		C.int(channels), C.int(chunksize)));
}

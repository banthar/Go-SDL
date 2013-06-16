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

#include "callback.h"
#include <stdio.h>
#include <pthread.h>
#include <string.h>
#include "_cgo_export.h"

static void SDLCALL callback(void *userdata, Uint8 *_stream, int len) {

	SDL_LockAudio();

	context context = { _stream, len };
	streamCallback(&context);
	
	SDL_UnlockAudio();
}

callback_t callback_getCallback() {
	return &callback;
}

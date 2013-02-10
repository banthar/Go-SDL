/*
 * Copyright: ⚛ <0xe2.0x9a.0x9b@gmail.com> 2010
 *
 * The contents of this file can be used freely,
 * except for usages in immoral contexts.
 */

#include "callback.h"
#include <stdio.h>
#include <pthread.h>
#include <string.h>
#include "_cgo_export.h"

// callbackFunc holds the callback library function.
// It is stored in a function pointer because C linkage
// does not work across packages.
static void(*callbackFunc)(void (*f)(void*), void*);

void setCallbackFunc(void *cb){ callbackFunc = cb; }

static void SDLCALL callback(void *userdata, Uint8 *_stream, int len) {

	SDL_LockAudio();

	context context = { _stream, len };
	callbackFunc(streamCallback, &context);
	
	SDL_UnlockAudio();
}

callback_t callback_getCallback() {
	return &callback;
}

/*
 * Copyright neagix - Feb 2013
 * Copyright: ⚛ <0xe2.0x9a.0x9b@gmail.com> 2010
 * 
 *
 * The contents of this file can be used freely,
 * except for usages in immoral contexts.
 * 
 */

#ifndef	__CALLBACK_H
#define	__CALLBACK_H

#include <SDL_audio.h>

typedef struct context context;
struct context {
	void *Stream;
	int NumBytes;
};

typedef void (SDLCALL *callback_t)(void *userdata, Uint8 *stream, int len);

extern callback_t callback_getCallback();

#endif //__CALLBACK_H

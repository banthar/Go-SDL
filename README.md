# Introduction

This is an improved version of 0xe2-0x9a-0x9b's [Go-SDL](https://github.com/0xe2-0x9a-0x9b/Go-SDL)
currently mantained by neagix.

The improvements/differences are:

* audio callback support
* downstreaming support

There is a nice and fully working PC speaker buzzer example in examples/callback.

# Known issues

The re-designed audio system supports only signed 16bit samples, but writing the others is as easy as a copy/paste.

# Installation

Make sure you have SDL, SDL-image, SDL-mixer and SDL-ttf (all in -dev version).

Installing libraries and examples:

    go get -v github.com/neagix/Go-SDL/sdl
    go get -v github.com/neagix/Go-SDL/sdl/audio


# Credits

Music to test SDL-mixer is by Kevin MacLeod.

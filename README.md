# Introduction

This is an improved version of Banthar's [Go-SDL](http://github.com/banthar/Go-SDL).

The improvements/differences are:

* SDL functions (except for SDL-mixer) can be safely called from concurrently
  running goroutines
* All SDL events are delivered via a Go channel
* Support for low-level SDL sound functions

* Can be installed in parallel to Banthar's Go-SDL
* The import paths are "github.com/0xe2-0x9a-0x9b/Go-SDL/..."


# Installation

Make sure you have SDL, SDL-image, SDL-mixer and SDL-ttf (all in -dev version).

Installing libraries and examples:

    go get -d -v  github.com/0xe2-0x9a-0x9b/Go-SDL
    go install -v github.com/0xe2-0x9a-0x9b/Go-SDL/...


# Credits

Music to test SDL-mixer is by Kevin MacLeod.

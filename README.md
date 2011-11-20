# Introduction

This is an improved version of Banthar's [Go-SDL](http://github.com/banthar/Go-SDL).

The improvements/differences are:

* SDL functions (except for SDL-mixer) can be safely called from multiple concurrent goroutines
* All SDL events are delivered via a Go channel
* Support for low-level SDL sound functions

* Can be installed in parallel to Banthar's Go-SDL
* The import path is "atom/sdl", instead of "sdl"


# Installation

Make sure you have SDL, SDL-image, SDL-mixer, and SDL-ttf (all in -dev version).

To install, run 'make' in the top-level directory.  If it fails to compile, try to run 'hg pull; hg update release' in $GOROOT and rebuild Go.


# Credits

Music to test SDL-mixer is by Kevin MacLeod.

Go-SDL
======

Go-SDL provides bindings for the [SDL][sdl], [SDL_image][sdl-image], [SDL_mixer][sdl-mixer], and [SDL_ttf][sdl-ttf] libraries.

Prerequisites
-------------

 * A version of [Go][go] compatible with weekly.2011-12-14
 * [SDL][sdl]
 * [SDL_image][sdl-image]
 * [SDL_mixer][sdl-mixer] (For *mixer* package only)
 * [SDL_ttf][sdl-ttf] (For *ttf* package only)

Installation
------------

To install all of the relevant libraries, use the following:

    goinstall github.com/banthar/Go-SDL/sdl
    goinstall github.com/banthar/Go-SDL/ttf
    goinstall github.com/banthar/Go-SDL/gfx
    goinstall github.com/banthar/Go-SDL/mixer

If you don't have write permission for GOPATH/GOROOT, you may need to run the previous command as root. If you get errors while trying to run it using sudo, it's possible that the GOROOT/GOOS/GOARCH/GOBIN variables are not available to the make command. You can try using '-E' to preserve the environment:

    sudo -E goinstall github.com/banthar/Go-SDL/sdl
    sudo -E goinstall github.com/banthar/Go-SDL/ttf
    sudo -E goinstall github.com/banthar/Go-SDL/gfx
    sudo -E goinstall github.com/banthar/Go-SDL/mixer

It's also possible to install just using make:

    make install

or

    sudo -E make install

Usage
-----

To import, use the following:

    import "github.com/banthar/Go-SDL/sdl"

Replace the final 'sdl' with the library that you want to import.

Credits
-------

 * [banthar](https://github.com/banthar)
 * Kevin MacLeod (*mixer* test music)

[go]: http://www.golang.org
[sdl]: http://www.libsdl.org
[sdl-image]: http://www.libsdl.org/projects/SDL_image/
[sdl-mixer]: http://www.libsdl.org/projects/SDL_mixer/
[sdl-ttf]: http://www.libsdl.org/projects/SDL_ttf/

<!--
    vim:ts=4 sw=4 et
-->

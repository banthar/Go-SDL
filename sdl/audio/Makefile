include $(GOROOT)/src/Make.inc

TARG=⚛sdl/audio

CGOFILES=audio.go
CGO_LDFLAGS=`pkg-config --libs sdl` -lrt
CGO_CFLAGS=`pkg-config --cflags sdl`
CGO_OFILES=callback.o

include $(GOROOT)/src/Make.pkg

_cgo_defun.c: callback.h
callback.o: callback.h

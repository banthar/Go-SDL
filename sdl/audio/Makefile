include $(GOROOT)/src/Make.inc

TARG=atom/sdl/audio

CGOFILES=audio.go
CGO_CFLAGS=`pkg-config --cflags sdl`
CGO_OFILES=callback.o

ifeq ($(GOOS),windows)
CGO_LDFLAGS=`pkg-config --libs sdl` -lpthread
else
CGO_LDFLAGS=`pkg-config --libs sdl` -lrt
endif

include $(GOROOT)/src/Make.pkg

_cgo_defun.c: callback.h
callback.o: callback.h

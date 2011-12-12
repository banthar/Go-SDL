include $(GOROOT)/src/Make.inc

TARG=atom/sdl/audio

CGOFILES=audio.go
CGO_OFILES=callback.o

ifeq ($(GOOS),windows)
CGO_LDFLAGS=-lpthread
else
CGO_LDFLAGS=-lrt
endif

include $(GOROOT)/src/Make.pkg

_cgo_defun.c: callback.h
callback.o: callback.h

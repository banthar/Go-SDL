include $(GOROOT)/src/Make.inc

TARG=⚛sdl/audio
CALLBACK_LIB=sdl_audio_callback.so

CGOFILES=audio.go
CGO_LDFLAGS=$(CALLBACK_LIB) `pkg-config --libs sdl` -lrt
CGO_CFLAGS=`pkg-config --cflags sdl`
CGO_DEPS=$(CALLBACK_LIB)

INSTALLFILES+=$(pkgdir)/$(CALLBACK_LIB)

include $(GOROOT)/src/Make.pkg

_cgo_defun.c: callback.h


_obj/$(TARG).a: $(CALLBACK_LIB)

callback.o: callback.c
	$(HOST_CC) $(_CGO_CFLAGS_$(GOARCH)) -fPIC -O2 -o $@ -c $(CGO_CFLAGS) $*.c

callback.o: callback.h

$(CALLBACK_LIB): callback.o
	$(HOST_CC) $(_CGO_CFLAGS_$(GOARCH)) -o $@ $^ $(_CGO_LDFLAGS_$(GOOS)) -lrt

$(pkgdir)/$(CALLBACK_LIB): $(CALLBACK_LIB)
	cp $^ $@

include $(GOROOT)/src/Make.inc

TARG=sdl/transform
CGOFILES=transform.go
CGO_OFILES=scale.${O}.o transform.o
CGO_LDFLAGS=`pkg-config --libs sdl` -lgcc_s
CGO_CFLAGS:=`pkg-config --cflags sdl`
include $(GOROOT)/src/Make.pkg

scale.${O}.o: scale.${O}.c
transform.o: transform.c

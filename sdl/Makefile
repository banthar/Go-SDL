
include $(GOROOT)/src/Make.inc

TARG=sdl

GOFILES:=constants.go structs.$(O).go

CGOFILES:=sdl.go

CGO_LDFLAGS:=-Lc:/msys/1.0/local/lib -lSDL -lSDL_image

CGO_CFLAGS:=-Ic:/msys/1.0/local/include

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

constants.go: constants.c
	godefs -g sdl constants.c > constants.go
	gofmt -w constants.go


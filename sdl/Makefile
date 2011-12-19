
include $(GOROOT)/src/Make.inc

TARG=github.com/banthar/Go-SDL/sdl

GOFILES:=constants.go structs_$(GOARCH).go

CGOFILES:=sdl.go

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

constants.go: constants.c
	godefs -g sdl constants.c > constants.go
	gofmt -w constants.go

fmt:
	gofmt -w $(wildcard *.go)

include $(GOROOT)/src/Make.inc

TARG=exp/gui/sdl
GOFILES=\
		gui.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $(wildcard *.go)
	gomake -C test fmt

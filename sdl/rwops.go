package sdl

import (
	"log"
	"unsafe"
)

// #cgo pkg-config: sdl
// #include <SDL.h>
import "C"

type RWops struct {
	cRWops *C.SDL_RWops

	gcBytes []byte // Prevents garbage collection of memory passed to RWFromMem
}

func (s *RWops) destroy() {
	s.cRWops = nil
	s.gcBytes = nil
}

// Creates an RWops from memory.
func RWFromMem(buf []byte) *RWops {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	p := C.SDL_RWFromMem(unsafe.Pointer(&buf[0]), C.int(len(buf)))
	var rwops RWops
	rwops.cRWops = (*C.SDL_RWops)(p)
	rwops.gcBytes = buf
	return &rwops

}

func (self *RWops) Free() {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	C.SDL_FreeRW(self.cRWops)
	self.cRWops = nil
	self.gcBytes = nil
}

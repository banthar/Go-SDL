/* 
A pure Go version of SDL_framerate
*/

package gfx

import (
	"time"
)

type FPSmanager struct {
	framecount uint32
	rateticks  float
	lastticks  uint64
	rate       uint32
}

func NewFramerate() *FPSmanager {
	return &FPSmanager{
		framecount: 0,
		rate:       FPS_DEFAULT,
		rateticks:  (1000.0 / float(FPS_DEFAULT)),
		lastticks:  uint64(time.Nanoseconds()) / 1e6,
	}
}

func (manager *FPSmanager) SetFramerate(rate uint32) {
	if rate >= FPS_LOWER_LIMIT && rate <= FPS_UPPER_LIMIT {
		manager.framecount = 0
		manager.rate = rate
		manager.rateticks = 1000.0 / float(rate)
	} else {
	}
}

func (manager *FPSmanager) GetFramerate() uint32 {
	return manager.rate
}

func (manager *FPSmanager) FramerateDelay() {
	var current_ticks, target_ticks, the_delay uint64

	// next frame
	manager.framecount++

	// get/calc ticks
	current_ticks = uint64(time.Nanoseconds()) / 1e6
	target_ticks = manager.lastticks + uint64(float(manager.framecount)*manager.rateticks)

	if current_ticks <= target_ticks {
		the_delay = target_ticks - current_ticks
		time.Sleep(int64(the_delay) * 1e6)
	} else {
		manager.framecount = 0
		manager.lastticks = uint64(time.Nanoseconds()) / 1e6
	}
}

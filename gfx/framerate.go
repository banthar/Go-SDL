/* 
A pure Go version of SDL_framerate
*/

package gfx

import (
	"sdl"
)

type FPSmanager struct {
<<<<<<< HEAD
  framecount uint32
  rateticks float64
  lastticks uint32
  rate uint32
}

func NewFramerate() *FPSmanager {
  return &FPSmanager{
    framecount: 0,
    rate: FPS_DEFAULT,
    rateticks: (1000.0 / float64(FPS_DEFAULT)),
    lastticks: sdl.GetTicks(),
  }
}

func (manager *FPSmanager) SetFramerate(rate uint32) {
  if rate >= FPS_LOWER_LIMIT && rate <= FPS_UPPER_LIMIT {
    manager.framecount = 0
    manager.rate = rate
    manager.rateticks = 1000.0 / float64(rate)
  } else {
  }
=======
	framecount uint32
	rateticks  float32
	lastticks  uint32
	rate       uint32
}

func NewFramerate() *FPSmanager {
	return &FPSmanager{
		framecount: 0,
		rate:       FPS_DEFAULT,
		rateticks:  (1000.0 / float32(FPS_DEFAULT)),
		lastticks:  sdl.GetTicks(),
	}
}

func (manager *FPSmanager) SetFramerate(rate uint32) {
	if rate >= FPS_LOWER_LIMIT && rate <= FPS_UPPER_LIMIT {
		manager.framecount = 0
		manager.rate = rate
		manager.rateticks = 1000.0 / float32(rate)
	} else {
	}
>>>>>>> 431cd07b5149e29b97c9315b845a1108ee3af468
}

func (manager *FPSmanager) GetFramerate() uint32 {
	return manager.rate
}

func (manager *FPSmanager) FramerateDelay() {
<<<<<<< HEAD
  var current_ticks, target_ticks, the_delay uint32

  // next frame
  manager.framecount++

  // get/calc ticks
  current_ticks = sdl.GetTicks()
  target_ticks = manager.lastticks + uint32(float64(manager.framecount) * manager.rateticks)

  if current_ticks <= target_ticks {
    the_delay = target_ticks - current_ticks
    sdl.Delay(the_delay)
  } else {
    manager.framecount = 0
    manager.lastticks = sdl.GetTicks()
  }
=======
	var current_ticks, target_ticks, the_delay uint32

	// next frame
	manager.framecount++

	// get/calc ticks
	current_ticks = sdl.GetTicks()
	target_ticks = manager.lastticks + uint32(float32(manager.framecount)*manager.rateticks)

	if current_ticks <= target_ticks {
		the_delay = target_ticks - current_ticks
		sdl.Delay(the_delay)
	} else {
		manager.framecount = 0
		manager.lastticks = sdl.GetTicks()
	}
>>>>>>> 431cd07b5149e29b97c9315b845a1108ee3af468
}

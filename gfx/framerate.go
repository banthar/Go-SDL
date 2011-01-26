/* 
A pure Go version of SDL_framerate
*/

package gfx

import (
  "sdl"
)

type FPSmanager struct {
  framecount uint32
  rateticks float32
  lastticks uint32
  rate uint32
}

func NewFramerate() *FPSmanager {
  return &FPSmanager{
    framecount: 0,
    rate: FPS_DEFAULT,
    rateticks: (1000.0 / float32(FPS_DEFAULT)),
    lastticks: sdl.GetTicks(),
  }
}

func (manager *FPSmanager) SetFramerate(rate uint32) {
  if rate >= FPS_LOWER_LIMIT && rate <= FPS_UPPER_LIMIT {
    manager.framecount = 0
    manager.rate = rate
    manager.rateticks = 1000.0 / float32(rate)
  } else {
  }
}

func (manager *FPSmanager) GetFramerate() uint32 {
  return manager.rate
}

func (manager *FPSmanager) FramerateDelay() {
  var current_ticks, target_ticks, the_delay uint32

  // next frame
  manager.framecount++

  // get/calc ticks
  current_ticks = sdl.GetTicks()
  target_ticks = manager.lastticks + uint32(float32(manager.framecount) * manager.rateticks)

  if current_ticks <= target_ticks {
    the_delay = target_ticks - current_ticks
    sdl.Delay(the_delay)
  } else {
    manager.framecount = 0
    manager.lastticks = sdl.GetTicks()
  }
}

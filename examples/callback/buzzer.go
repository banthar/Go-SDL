/*
 * Copyright neagix 2013
 * This sample is part of langwar project
 * https://github.com/neagix/langwar
 * 
 * Licensed under GNU/GPL v2
 */

package main

import "github.com/neagix/Go-SDL/sdl/audio"
import "math"
import "time"

type BeepObject struct {
	Freq        float64
	SamplesLeft int
}

var Play chan BeepObject

const AMPLITUDE = 28000
const FREQUENCY = 44100
const AUDIO_SAMPLES = 2048

func Beep(freq, duration int) {
	bo := BeepObject{
		Freq:        float64(freq),
		SamplesLeft: duration * FREQUENCY / 1000}

	Play <- bo
}

func Init() bool {

	desiredSpec := audio.AudioSpec{
		Freq:     FREQUENCY,
		Format:   audio.AUDIO_S16SYS,
		Channels: 1,
		Samples:  AUDIO_SAMPLES,
	}
	var obtainedSpec audio.AudioSpec

	if audio.OpenAudio(&desiredSpec, &obtainedSpec) != 0 {
		return false
	}

	Play = make(chan BeepObject)

	// start the playback queue processor
	go func() {
		for {
			// pick next beep object
			bo := <-Play

			stream := make([]int16, bo.SamplesLeft)

			v := float64(0)
			for i := 0; i < bo.SamplesLeft; i++ {
				stream[i] = int16(AMPLITUDE * math.Sin(v*2*math.Pi/FREQUENCY))
				v += bo.Freq
			}

			audio.SendAudio_int16(stream)
		}
	}()

	return true
}

// to be deferred after corresponding Init()
func Quit() {
	defer audio.CloseAudio()
}

// song by niniel1 as found at http://gendou.com/t/20439
func PlaySong1() {
	Beep(349, 400)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(349, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(349, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(262, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(349, 400)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 267)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(349, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(392, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(440, 133)
	time.Sleep(time.Millisecond * 33)
	Beep(349, 533)
	time.Sleep(time.Millisecond * 33)
}

// find out more songs at https://github.com/binarypearl/beepbeep
func PlaySong3() {
	Beep(784, 100)
	Beep(784, 100)
	Beep(784, 100)
	time.Sleep(time.Millisecond * 100)
	Beep(784, 600)
	Beep(622, 600)
	Beep(698, 600)
	Beep(784, 200)
	time.Sleep(time.Millisecond * 200)
	Beep(698, 200)
	Beep(784, 800)

}

func main() {
	if !Init() {
		return
	}
	defer Quit()

	// start playback, yeah!
	audio.PauseAudio(false)

	PlaySong1()

	// it's a silence :)
	Beep(0, 400)

	PlaySong3()
}

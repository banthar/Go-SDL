package main

import (
	"github.com/neagix/Go-SDL/sdl"
	"github.com/neagix/Go-SDL/sdl/audio"
	"github.com/neagix/Go-SDL/sound"
	"log"
	"time"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)

	var resourcePath string
	{
		GOPATH := os.Getenv("GOPATH")
		if GOPATH == "" {
			log.Fatal("No such environment variable: GOPATH")
		}
		for _, gopath := range strings.Split(GOPATH, ":") {
			a := gopath + "/src/github.com/neagix/Go-SDL/examples/completeTest"
			_, err := os.Stat(a)
			if err == nil {
				resourcePath = a
				break
			}
		}
		if resourcePath == "" {
			log.Fatal("Failed to find resource directory")
		}
	}

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		log.Fatal(sdl.GetError())
	}

	sound.Init()

	desiredSpec := audio.AudioSpec{Freq: 44100, Format: audio.AUDIO_S16SYS, Channels: 1, Samples: 4096}
	var obtainedSpec audio.AudioSpec

	if audio.OpenAudio(&desiredSpec, &obtainedSpec) != 0 {
		log.Fatal(sdl.GetError())
	}

	fmt := sound.AudioInfo{obtainedSpec.Format, obtainedSpec.Channels, uint32(obtainedSpec.Freq)}
	/* note: on my machine i always get 2 channels, despite only requesting one */

	sdl.EnableUNICODE(1)

	sample := sound.NewSampleFromFile(resourcePath+"/test.ogg", &fmt, 1048756)

	sample.DecodeAll()
	buf := sample.Buffer_int16()
	/* this decodes the entire file at once and loads it into memory. sample.Decode()
	will decode in chunks of ~1mb, since that's the buffer size I requested above. */

	audio.PauseAudio(false)

	audio.SendAudio_int16(buf)
	/* note: this sends the entire buffer to the audio system at once, which
	is a stupid thing to do in practice because the audio callback will block while
	it copies the samples to its internal tail buffer, and won't be able to fill the
	current audio frame in a timely manner. this will cause many underruns.
	A better approach is to call SendAudio in a loop with small chunks of the
	buffer at a time. */

	time.Sleep(time.Second * 45)
	/* we better stick around or we'll exit before making any noise! */
}

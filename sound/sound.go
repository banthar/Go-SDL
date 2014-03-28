/*
A binding of SDL_sound.

Currently NewSampleFromFile() is the only way to create a Sample struct, the
core focus of the API. From there you can call Sample.Decode()/DecodeAll()
followed by Sample.Buffer_int16() to get at the decoded sound samples.

Signed 16-bit ints is the only supported audio format at the moment.
*/
package sound

// #cgo pkg-config: sdl
// #include <SDL_sound.h>
// #cgo LDFLAGS: -lSDL_sound
import "C"
import "unsafe"

import "errors"
import "sync"
import "fmt"
import "github.com/neagix/Go-SDL/sdl/audio"

type AudioInfo struct {
	Format   uint16
	Channels uint8
	Rate     uint32
}

type DecoderInfo struct {
	Extensions  []string // a list of valid extensions
	Description string   // human readable description
	Author      string
	Url         string
}

type Sample struct {
	sync.Mutex
	csample *C.Sound_Sample
	Decoder *DecoderInfo
	Desired *AudioInfo
	Actual  *AudioInfo

	nbytes uint32 // number of bytes read in last Decode()
}

func GetError() error {
	errstr := C.GoString(C.Sound_GetError())
	C.Sound_ClearError()
	return errors.New(errstr)
}

func Init() int {
	i, _ := C.Sound_Init()
	return int(i)
}

func Quit() int {
	i, _ := C.Sound_Quit()
	return int(i)
}

func fromCDecoderInfo(cinfo *C.Sound_DecoderInfo) *DecoderInfo {
	if cinfo == nil {
		return nil
	}
	info := DecoderInfo{}
	extptr := uintptr(unsafe.Pointer((*cinfo).extensions))
	for {
		ext := (**C.char)(unsafe.Pointer(extptr))
		if *ext == nil {
			break
		}
		info.Extensions = append(info.Extensions, C.GoString(*ext))
		extptr += unsafe.Sizeof(extptr)
	}
	info.Description = C.GoString((*cinfo).description)
	info.Author = C.GoString((*cinfo).author)
	info.Url = C.GoString((*cinfo).url)
	return &info
}

func AvailableDecoders() []DecoderInfo {
	cinfo, _ := C.Sound_AvailableDecoders()
	infos := make([]DecoderInfo, 0, 16)
	infptr := uintptr(unsafe.Pointer(cinfo))
	for {
		cinfo = (**C.Sound_DecoderInfo)(unsafe.Pointer(infptr))
		if *cinfo == nil {
			break
		}
		infos = append(infos, *fromCDecoderInfo(*cinfo))
		infptr += unsafe.Sizeof(infptr)
	}
	return infos
}

func fromCAudioInfo(cinfo *C.Sound_AudioInfo) *AudioInfo {
	if cinfo == nil {
		return nil
	}
	return &AudioInfo{uint16(cinfo.format), uint8(cinfo.channels), uint32(cinfo.rate)}
}

func cAudioInfo(info *AudioInfo) *C.Sound_AudioInfo {
	if info == nil {
		return nil
	}
	cinfo := new(C.Sound_AudioInfo)
	cinfo.format = C.Uint16(info.Format)
	cinfo.channels = C.Uint8(info.Channels)
	cinfo.rate = C.Uint32(info.Rate)
	return cinfo
}

func NewSampleFromFile(filename string, desired *AudioInfo, size uint) (*Sample, error) {
	sample := new(Sample)
	cfile := C.CString(filename)
	defer C.free(unsafe.Pointer(cfile))
	sample.csample = C.Sound_NewSampleFromFile(cfile, cAudioInfo(desired), C.Uint32(size))
	if sample.csample == nil {
		return nil, GetError()
	}
	sample.Decoder = fromCDecoderInfo(sample.csample.decoder)
	sample.Desired = fromCAudioInfo(&sample.csample.desired)
	sample.Actual = fromCAudioInfo(&sample.csample.actual)
	return sample, nil
}

func FreeSample(sample Sample) {
	C.Sound_FreeSample(sample.csample)
}

func (sample *Sample) SetBufferSize(size uint32) int {
	ret := C.Sound_SetBufferSize(sample.csample, C.Uint32(size))
	return int(ret)
}

/* Decodes as many samples as will fit in the buffer.
 * Returns the number of BYTES read (zero at EOF). */
func (sample *Sample) Decode() uint32 {
	sample.Lock()
	defer sample.Unlock()
	sample.nbytes = uint32(C.Sound_Decode(sample.csample))
	return sample.nbytes
}

func (sample *Sample) DecodeAll() uint32 {
	sample.Lock()
	defer sample.Unlock()
	sample.nbytes = uint32(C.Sound_DecodeAll(sample.csample))
	return sample.nbytes
}

func (sample *Sample) Rewind() int {
	ret := C.Sound_Rewind(sample.csample)
	return int(ret)
}

func (sample *Sample) Seek(ms uint32) int {
	ret := C.Sound_Seek(sample.csample, C.Uint32(ms))
	return int(ret)
}

func (sample *Sample) Flags() uint {
	return uint(sample.csample.flags)
}

func (sample *Sample) Buffer_int16() []int16 {
	sample.Lock()
	defer sample.Unlock()
	if sample.Desired.Format != audio.AUDIO_S16SYS {
		panic(fmt.Sprintf("wrong format requested %d", sample.Desired.Format))
	}

	buf := make([]int16, int(sample.nbytes)/2)
	for i := 0; i < len(buf); i++ {
		buf[i] = int16(*((*C.int16_t)(unsafe.Pointer((uintptr(sample.csample.buffer) + uintptr(i*2))))))
	}
	return buf
}

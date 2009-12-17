package mixer

//TODO: Autogenerate these values and add rest of constants
const (
	AUDIO_U8          = 0x0008
	AUDIO_S8          = 0x8008
	AUDIO_U16LSB      = 0x0010
	AUDIO_S16LSB      = 0x8010
	AUDIO_U16MSB      = 0x1010
	AUDIO_S16MSB      = 0x9010
	AUDIO_U16         = 0x0010
	AUDIO_S16         = 0x8010
	DEFAULT_FREQUENCY = 22050
	DEFAULT_FORMAT    = 0x8010
	DEFAULT_CHANNELS  = 2
	MAX_VOLUME        = 128
)

const (
	MUS_NONE = iota
	MUS_CMD
	MUS_WAV
	MUS_MOD
	MUS_MID
	MUS_OGG
	MUS_MP3
)

const (
	NO_FADING = iota
	FADING_OUT
	FADING_IN
)

package sound

const (
	SAMPLEFLAG_NONE = 0
	SAMPLEFLAG_CANSEEK = 1
	SAMPLEFLAG_EOF = 1 << 29
	SAMPLEFLAG_ERROR = 1 << 30	// unrecoverable error
	SAMPLEFLAG_EAGAIN = 1 << 31	// function would block or temp error
)

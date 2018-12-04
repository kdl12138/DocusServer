package Log

const (
	DebugLevel = 1
	InfoLevel  = 2
	WarnLevel  = 4
	ErrorLevel = 8
	PanicLevel = 16
)

type Level uint8

func (Level Level) IsAllow(lvl Level) bool {
	return Level <= lvl
}

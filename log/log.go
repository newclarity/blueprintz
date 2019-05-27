package log

import "github.com/gearboxworks/go-status"

type (
	Msg = status.Msg
)

func NewLogger() *Logger {
	return &Logger{}
}

var NilL = (*Logger)(nil)
var _ status.MsgLogger = NilL

type Logger struct {
	status.L
}

func (me *Logger) Debug(msg Msg) {
	//fmt.Printf("[DEBUG] %s\n", msg)
}

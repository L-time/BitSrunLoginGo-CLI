package terminal

import "github.com/fatih/color"

type OutPut struct {
	color.Color
}

func (o *OutPut) SwitchColor(b bool) {
	switch b {
	case true:
		o.EnableColor()
	case false:
		o.DisableColor()
	}
}

func (o *OutPut) Info(s string, a ...interface{}) {
	color.Green(s)
}

func (o *OutPut) Warn(s string, a ...interface{}) {
	color.Yellow(s)
}

func (o *OutPut) ErrorWithError(f string, err error, a ...interface{}) {
	color.Red(f, err, a)
}

func (o *OutPut) ErrorWithString(s string, a ...interface{}) {
	color.Red(s, a)
}

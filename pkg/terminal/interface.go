package terminal

// Terminal 执行期的彩色终端界面接口，非在命令解释期使用！
type Terminal interface {
	SwitchColor(b bool)
	Info(s string, a ...interface{})
	Warn(s string, a ...interface{})
	ErrorWithError(f string, err error, a ...interface{})
	ErrorWithString(s string, a ...interface{})
}

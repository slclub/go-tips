package logf

/**  interface of log
 *   Logger is simple and small
 * 	 Convenient injection into your own log system
 */
type Logger interface {
	Printf(string, ...any)
	Print(...any)
}

var _log Logger

func Log() Logger {
	return _log
}

func LogOf(lg Logger) {
	if lg == nil {
		return
	}
	_log = lg
}

func Printf(format string, args ...any) {
	_log.Printf(format, args...)
}

func Print(args ...any) {
	_log.Print(args...)
}

func init() {
	LogOf(New())
}

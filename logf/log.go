package logf

import "fmt"

const (
	ECHO_STD = 1
)

type loger struct {
	echo int
}

func New() *loger {
	return &loger{
		echo: ECHO_STD,
	}
}

func (this *loger) Printf(format string, args ...any) {
	if this.echo&ECHO_STD > 0 {
		this.fmtPrintf(format, args...)
	}
}

func (this *loger) Print(args ...any) {
	if this.echo&ECHO_STD > 0 {
		this.fmtPrint(args...)
	}
}

func (this *loger) fmtPrintf(format string, args ...any) {
	fmt.Printf(format, args...)
	fmt.Print("\n")
}
func (this *loger) fmtPrint(args ...any) {
	fmt.Println(args...)
}

var _ Logger = &loger{}

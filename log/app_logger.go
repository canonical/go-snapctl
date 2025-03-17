package log

import (
	"fmt"
	"os"
)

type appLogger struct {
	prefix string
	debug  bool
}

func newAppLogger(label string, debug bool) (*appLogger, error) {
	return &appLogger{
		prefix: label + ": ",
		debug:  debug,
	}, nil
}

func (l *appLogger) Print(a ...any) {
	l.stdout(a...)
}

func (l *appLogger) Printf(format string, a ...any) {
	l.Print(fmt.Sprintf(format, a...))
}

func (l *appLogger) Debug(a ...any) {
	if l.debug {
		l.stdout(a...)
	}
}

func (l *appLogger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *appLogger) Error(a ...any) {
	l.stderr(a...)
}

func (l *appLogger) Errorf(format string, a ...any) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *appLogger) Fatal(a ...any) {
	l.Error(a...)
	os.Exit(1)
}

func (l *appLogger) Fatalf(format string, a ...any) {
	l.Errorf(format, a...)
	os.Exit(1)
}

// stdout writes the given input to standard output.
// It formats similar to [fmt.Sprint], adds the prefix, and appends a newline
// The newline is added for consistency with the syslog writer, and for better portability with Go log package
func (l *appLogger) stdout(a ...any) {
	fmt.Fprintf(os.Stdout, l.prefix+"%s\n", a...)
}

// stdout writes the given input to standard error.
// It formats similar to [fmt.Sprint], adds the prefix, and appends a newline
func (l *appLogger) stderr(a ...any) {
	fmt.Fprintf(os.Stderr, l.prefix+"%s\n", a...)
}

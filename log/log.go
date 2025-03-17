package log

var globalLogger logger

type logger interface {
	Print(...any)
	Printf(string, ...any)
	Debug(...any)
	Debugf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
	Fatal(...any)
	Fatalf(string, ...any)
}

// Print writes a normal log message
// It formats similar to [fmt.Sprint]
func Print(a ...any) {
	globalLogger.Print(a...)
}

// Printf writes a normal log message
// It formats similar to [fmt.Sprintf]
func Printf(format string, a ...any) {
	globalLogger.Printf(format, a...)
}

// Debug writes a debug log message
// It formats similar to [fmt.Sprint]
func Debug(a ...any) {
	globalLogger.Debug(a...)
}

// Debugf writes a debug log message
// It formats similar to [fmt.Sprintf]
func Debugf(format string, a ...any) {
	globalLogger.Debugf(format, a...)
}

// Error writes an error message
// It formats similar to [fmt.Sprint]
func Error(a ...any) {
	globalLogger.Error(a...)
}

// Errorf writes an error message
// It formats similar to [fmt.Sprintf]
func Errorf(format string, a ...any) {
	globalLogger.Errorf(format, a...)
}

// Fatal writes an error message and exits
// It formats similar to [fmt.Sprint]
func Fatal(a ...any) {
	globalLogger.Fatal(a...)
}

// Fatalf writes an error message and exits
// It formats similar to [fmt.Sprintf]
func Fatalf(format string, a ...any) {
	globalLogger.Fatalf(format, a...)
}

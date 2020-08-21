package scurvy

// Logger is the interface that should be met by a logging implementation
// in order for it to be useful to scurvy
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

// NullLogger is a default null logger to be used by various parts of the system
// if there is no other logger defined
type NullLogger struct{}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

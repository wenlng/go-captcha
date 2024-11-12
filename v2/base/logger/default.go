package logger

var defaultLog = New()

// Infof .
func Infof(format string, v ...interface{}) {
	defaultLog.Infof(format, v...)
}

// Errorf .
func Errorf(format string, v ...interface{}) {
	defaultLog.Errorf(format, v...)
}

// Warnf .
func Warnf(format string, v ...interface{}) {
	defaultLog.Warnf(format, v...)
}

// Debugf .
func Debugf(format string, v ...interface{}) {
	defaultLog.Debugf(format, v...)
}

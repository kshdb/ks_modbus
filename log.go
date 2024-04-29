package modbus

import (
	"os"
	"sync/atomic"

	"log"
)

// 内部调试实现.
type logger struct {
	provider LogProvider
	// has log output enabled,
	// 1: enable
	// 0: disable
	has uint32
}

// newLogger new logger with prefix.
func newLogger(prefix string) logger {
	return logger{
		provider: defaultLogger{log.New(os.Stdout, prefix, log.LstdFlags)},
		has:      0,
	}
}

// LogMode set enable or disable log output when you has set logger.
func (sf *logger) LogMode(enable bool) {
	if enable {
		atomic.StoreUint32(&sf.has, 1)
	} else {
		atomic.StoreUint32(&sf.has, 0)
	}
}

// setLogProvider overwrite log provider.
func (sf *logger) setLogProvider(p LogProvider) {
	if p != nil {
		sf.provider = p
	}
}

// Error Log ERROR level message.
func (sf logger) Error(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.has) == 1 {
		sf.provider.Error(format, v...)
	}
}

// Debug Log DEBUG level message.
func (sf logger) Debug(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.has) == 1 {
		sf.provider.Debug(format, v...)
	}
}

// default log.
type defaultLogger struct {
	*log.Logger
}

// check implement LogProvider interface.
var _ LogProvider = (*defaultLogger)(nil)

// Error Log ERROR level message.
func (sf defaultLogger) Error(format string, v ...interface{}) {
	sf.Printf("[E]: "+format, v...)
}

// Debug Log DEBUG level message.
func (sf defaultLogger) Debug(format string, v ...interface{}) {
	sf.Printf("[D]: "+format, v...)
}

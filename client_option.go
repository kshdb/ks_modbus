package modbus

import (
	"time"

	"github.com/goburrow/serial"
)

// ClientProviderOption client provider option for user.
type ClientProviderOption func(ClientProvider)

// WithLogProvider set logger provider.
func WithLogProvider(provider LogProvider) ClientProviderOption {
	return func(p ClientProvider) {
		p.setLogProvider(provider)
	}
}

// WithEnableLogger enable log output when you has set logger.
func WithEnableLogger() ClientProviderOption {
	return func(p ClientProvider) {
		p.LogMode(true)
	}
}

// WithAutoReconnect set auto reconnect count.
// if cnt == 0, disable auto reconnect
// if cnt > 0 ,enable auto reconnect,but max 6.
func WithAutoReconnect(cnt byte) ClientProviderOption {
	return func(p ClientProvider) {
		p.SetAutoReconnect(cnt)
	}
}

// WithSerialConfig set serial config, only valid on serial.
func WithSerialConfig(config serial.Config) ClientProviderOption {
	return func(p ClientProvider) {
		p.setSerialConfig(config)
	}
}

// WithTCPTimeout set tcp Connect & Read timeout, only valid on TCP.
func WithTCPTimeout(t time.Duration) ClientProviderOption {
	return func(p ClientProvider) {
		p.setTCPTimeout(t)
	}
}

// Package config provides a struct to store the applications config
package config

import (
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

// AppConfig stores all the config values for our application
var AppConfig struct {
	OIDC    echojwtx.AuthConfig
	CRDB    crdbx.Config
	Logging loggingx.Config
	Server  echox.Config
	Tracing otelx.Config
}

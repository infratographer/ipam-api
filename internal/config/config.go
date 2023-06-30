// Package config provides a struct to store the applications config
package config

import (
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

// AppConfig stores all the config values for our application
var AppConfig struct {
	Events      EventsConfig
	OIDC        echojwtx.AuthConfig
	CRDB        crdbx.Config
	Logging     loggingx.Config
	Permissions permissions.Config
	Server      echox.Config
	Tracing     otelx.Config
}

// EventsConfig stores the configuration for an event publisher
type EventsConfig struct {
	Publisher events.PublisherConfig
}

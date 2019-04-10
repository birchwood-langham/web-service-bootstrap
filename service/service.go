package service

import (
	"gitlab.com/birchwoodlangham/go-web-service-application.git/api"
)

// Properties is a structure for providing usage and descriptions to the root command
type Properties struct {
	Usage            string
	ShortDescription string
	LongDescription  string
}

// NewProperties creates a new set of properties for the root command
func NewProperties(usage string, shortDescription string, longDescription string) *Properties {
	return &Properties{
		usage,
		shortDescription,
		longDescription,
	}
}

// Configuration provides the configuration for a service
type Configuration struct {
	Properties       *Properties
	InitializeRoutes func(*api.Server)
	Cleanup          func() error
}

// NewConfiguration creates a new service configuration
func NewConfiguration(props *Properties, routes func(*api.Server), cleanup func() error) *Configuration {
	return &Configuration{
		props,
		routes,
		cleanup,
	}
}

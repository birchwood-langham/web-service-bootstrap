package service

import "gitlab.com/birchwoodlangham/go-web-service-application.git/api"

type Application interface {
	Init() error
	InitializeRoutes(server *api.Server)
	Cleanup() error
	Properties() *Properties
}

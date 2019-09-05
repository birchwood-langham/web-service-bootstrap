package service

import "github.com/birchwood-langham/web-service-bootstrap/api"

type Application interface {
	Init() error
	InitializeRoutes(server *api.Server)
	Cleanup() error
	Properties() Properties
}

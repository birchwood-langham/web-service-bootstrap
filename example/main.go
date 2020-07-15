package main

import (
	"net/http"

	"github.com/birchwood-langham/web-service-bootstrap/api"
	"github.com/birchwood-langham/web-service-bootstrap/cmd"
	"github.com/birchwood-langham/web-service-bootstrap/logger"
	"github.com/birchwood-langham/web-service-bootstrap/service"
	"go.uber.org/zap"
)

type MyApp struct {
	log *zap.Logger
}

// Init performs any initialization that is required for my application
func (a *MyApp) Init() (err error) {
	a.log = logger.New(logger.ApplicationLogLevel()).With(zap.String("app", "MyApp"))
	return
}

// initialiseRoutes allows you to define the routes required for the service
// and the handlers for each route
func (a *MyApp) InitializeRoutes(s *api.Server) {
	s.Router.HandleFunc("/hello", a.hello).Methods("GET")
}

// Cleanup is called to cleanup the service before it shuts down, for example if you need
// to perform a controlled shut down and ensure all processes have completed before terminating
// the application, you would implement it here
func (a *MyApp) Cleanup() error {
	return nil
}

func (a *MyApp) Properties() service.Properties {
	return service.NewProperties("usage", "short description", "A long detailed description")
}

// This is the obligatory hello world example implementing a Hello World service with this library
func (a *MyApp) hello(w http.ResponseWriter, r *http.Request) {
	a.log.Info("Received a request to say hello", zap.String("context", "hello"), zap.Int("test", 10), zap.String("version", "0.1.0"))
	api.RespondWithJSON(w, http.StatusOK, "Hello, World!")
}

func main() {
	cmd.Execute(&MyApp{})
}

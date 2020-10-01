package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"

	"github.com/birchwood-langham/web-service-bootstrap/api"
	"github.com/birchwood-langham/web-service-bootstrap/cmd"
	"github.com/birchwood-langham/web-service-bootstrap/logger"
	"github.com/birchwood-langham/web-service-bootstrap/service"

	"go.uber.org/zap"
)

type MyApp struct {
	logFile *lumberjack.Logger
	log     *zap.Logger
}

// Init performs any initialization that is required for my application
func (a *MyApp) Init() (err error) {
	a.logFile = logger.DefaultLumberjackLogger()
	a.log = logger.New(logger.ApplicationLogLevel(), a.logFile).With(zap.String("app", "MyApp"))
	a.log.Debug("Initializing MyApp")
	return
}

// initialiseRoutes allows you to define the routes required for the service
// and the handlers for each route
func (a *MyApp) InitializeRoutes(s *api.Server) {
	a.log.Debug("MyApp Initializing Routes")
	s.Router.HandleFunc("/hello", a.hello).Methods("GET")
}

// Cleanup is called to cleanup the service before it shuts down, for example if you need
// to perform a controlled shut down and ensure all processes have completed before terminating
// the application, you would implement it here
func (a *MyApp) Cleanup() error {
	a.log.Debug("MyApp Cleaning up")
	_ = a.log.Sync()

	return a.logFile.Close()
}

func (a *MyApp) Properties() service.Properties {
	return service.NewProperties("usage", "short description", "A long detailed description")
}

// This is the obligatory hello world example implementing a Hello World service with this library
func (a *MyApp) hello(w http.ResponseWriter, _ *http.Request) {
	a.log.Info("Received a request to say hello", zap.String("context", "hello"), zap.Int("test", 10), zap.String("version", "0.1.0"))
	api.RespondWithJSON(w, http.StatusOK, "Hello, World!")
}

func main() {
	cmd.Execute(&MyApp{})
}

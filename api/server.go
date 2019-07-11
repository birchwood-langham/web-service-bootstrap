package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"gitlab.com/birchwoodlangham/go-web-service-application.git/config"
)

// ServerMessage represents fixed messages processes can send to the server process
type ServerMessage uint16

// Stop is the message to route to the server process when we want the server to terminate
const Stop ServerMessage = 0

// Server represents the Http Server we are creating to provide the web service we are building
type Server struct {
	Router         *mux.Router
	messageChannel chan ServerMessage
	host           string
	port           int
	server         *http.Server
}

// New creates a new api.Server instance running on the given host and port
func New(hostname string, port int, messageChannel chan ServerMessage) *Server {
	return &Server{
		host:           hostname,
		port:           port,
		messageChannel: messageChannel,
	}
}

// Initialize sets up the routes you want for your API server
func (s *Server) Initialize(initializeRoutes func(*Server)) {
	s.Router = mux.NewRouter()
	initializeRoutes(s)
}

// RespondWithError wraps an error message as a JSON structure and returns it as a Http Response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON wraps your payload into JSON structure and returns it as a Http Response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Errorf("Could not write response: %v", err)
	}
}

// Run launches you server
func (s *Server) Run() {
	writeTimeout := config.DefaultWriteTimeout
	readTimeout := config.DefaultReadTimeout
	idleTimeout := config.DefaultIdleTimeout

	if viper.IsSet(config.ServiceWriteTimeoutKey) {
		writeTimeout = viper.GetInt(config.ServiceWriteTimeoutKey)
	}

	if viper.IsSet(config.ServiceReadTimeoutKey) {
		readTimeout = viper.GetInt(config.ServiceReadTimeoutKey)
	}

	if viper.IsSet(config.ServiceIdleTimeoutKey) {
		idleTimeout = viper.GetInt(config.ServiceIdleTimeoutKey)
	}

	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.host, s.port),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		IdleTimeout:  time.Second * time.Duration(idleTimeout),
		Handler:      s.Router,
	}

	if err := s.server.ListenAndServe(); err != nil {
		serviceName := "Unspecified"

		if viper.IsSet(config.ServiceNameKey) {
			serviceName = viper.GetString(config.ServiceNameKey)
		}

		log.Errorf("Could not start %s service: %v\n", serviceName, err)
		// controlled stop by sending a stop message to the main thread
		s.messageChannel <- Stop
	}
}

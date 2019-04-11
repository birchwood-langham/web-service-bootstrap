# Birchwood Langham Go Web Application Quick Starter Library

This library provides a basic structure for creating a web application using Go. It leverages a few common Go libraries to provide a quick start bootstrap for creating a web application in Go.

## Third party libraries

The third party libraries we are using are:

| Library | Description | Link |
| ------- | ----------- | ---- |
| Cobra | Library for creating CLI applications | [https://github.com/spf13/cobra](https://github.com/spf13/cobra) |
| Viper | Library for application configuration | [https://github.com/spf13/viper](https://github.com/spf13/viper) |
| Gorilla Mux | Library for providing URL routing and dispatch | [https://github.com/gorilla/mux](https://github.com/gorilla/mux) |
| Logrus | Library for Structured, pluggable logging | [github.com/sirupsen/logrus](github.com/sirupsen/logrus) |

## Pre-requisites

This library requires Go version 1.11+ as it leverages Go Modules

## Usage

To use this library, create a new project and initialize the Go module

```bash
mkdir my-go-webapp
cd my-go-webapp
go mod init my-go-webapp
```

Use go get to add this project as a dependency to your Go module

```bash
go get gitlab.com/birchwoodlangham/go-web-service-application.git
```

> *__NOTE:__* Due to known issues with GitLab, you will need to include the .git at the end of the project name otherwise you won't be able to pull the project properly using go get.

To create and start your application, you simply need to implement the service.Application interface and pass you Application to the cmd.Execute() function to launch your application.

```go
// main.go
package main

import (
  "fmt"
  "net/http"

  "gitlab.com/birchwoodlangham/go-web-service-application.git/service"
  "gitlab.com/birchwoodlangham/go-web-service-application.git/cmd"
  "github.com/gorilla/mux"
)

type MyApp struct {}

// Init performs any initialization that is required for my application
func (a *MyApp) Init() { }

// initialiseRoutes allows you to define the routes required for the service
// and the handlers for each route
func (a *MyApp) InitializeRoutes(s *api.Server) {
  s.Router.HandleFunc("/hello", hello).Methods("GET")
}

// Cleanup is called to cleanup the service before it shuts down, for example if you need
// to perform a controlled shut down and ensure all processes have completed before terminating
// the application, you would implement it here
func (a *MyApp) Cleanup() error {
  return nil
}

func (a *MyApp) Properties() *service.Properties {
	return service.NewProperties("usage", "short description", "A long detailed description")
} 

// This is the obligatory hello world example implementing a Hello World service with this library
func hello(w http.ResponseWriter, r *http.Request) {
  api.RespondWithJSON(w, http.StatusOK, "Hello, World!")
}

func main() {
  cmd.Execute(MyApp{})
}
```

The application will look for an application.yaml file containing the properties it needs to run your application.

```yaml
# application.yaml
version: 0.0.1
service:
  name: my-go-webapp
  host: localhost
  port: 8989
  write-timeout-seconds: 20
  read-timeout-seconds: 20
  idle-timeout-seconds: 60
  api-command-buffer: 100
log-file-path: ./log/trend-risk.log
log-level: DEBUG
```

You can add additional configuration settings into this application.yaml file and they will be loaded and accessible via viper.

To add your own CLI commands, you can just create a command, and add them before calling the `cmd.Execute()` function. For example:

```go
package main

// omitted for clarity

var helloCmd = &cobra.Command{
    Use: "hello",
    Short: "Says hello",
    Long: "The obligatory hello world function",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, World!")
    },
}

func main() {
    cmd.AddCommand(helloCmd)
    cmd.Execute(MyApp{})
}
```


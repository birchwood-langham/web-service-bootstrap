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
go get github.com/birchwood-langham/go-web-service-application
```

To create and start your application, you simply need to implement the service.Application interface and pass you Application to the cmd.Execute() function to launch your application.

```go
// main.go
package main

import (
  "fmt"
  "net/http"
  "database/sql"

  "github.com/birchwood-langham/go-web-service-application/service"
  "github.com/birchwood-langham/go-web-service-application/cmd"
  "github.com/gorilla/mux"
)

type MyApp struct {
  db *sql.DB
}

func New() *MyApp {
  // At this point in the application execution, viper will not have been initialized
  // and you will not be able to read properties from your configuration file
  return &MyApp{}
}

// Init performs any initialization that is required for my application
func (a *MyApp) Init() (err error) { 
  // Once the application has been started, viper will have been configured, and Init is called to 
  // initialize whatever you need for your application

  // for example if the application needs access to a database, we can initialize it here before anything
  // else happens
  a.db, err = sql.Open(...) // you can use properties from Viper now to help initialize your application
  return 
}

// initialiseRoutes allows you to define the routes required for the service
// and the handlers for each route, you can define these routes using the methods defined by
// gorilla mux
func (a *MyApp) InitializeRoutes(s *api.Server) {
  s.Router.HandleFunc("/hello", hello).Methods("GET")
}

// Cleanup is called to cleanup the service before it shuts down, for example if you need
// to perform a controlled shut down and ensure all processes have completed before terminating
// the application, you would implement it here
func (a *MyApp) Cleanup() error {
  return nil
}

// properties provide the short, lomg, and usage information to be displayed by the application 
// if you pass --help on the command line
func (a *MyApp) Properties() *service.Properties {
	return service.NewProperties("usage", "short description", "A long detailed description")
} 

// This is the obligatory hello world example implementing a Hello World service with this library
func hello(w http.ResponseWriter, r *http.Request) {
  api.RespondWithJSON(w, http.StatusOK, "Hello, World!")
}

func main() {
  cmd.Execute(New())
}
```

The application will look for an application.yaml file containing the properties it needs to run your application. It will search the current directory and for a conf sub-directory for the application.yaml file.

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
    cmd.Execute(New())
}
```

## Configuration

Inspired by go-micro config library, I have added some wrapper functions around the viper Get functions to allow for default values to be substituted in if the data is not present in the config file.

### Example

Instead of doing something like this:

```go

    serviceName := "Default Service Name"
    
    if viper.IsSet("service.name") {
    	serviceName = viper.Get("service.name")
    }

```

We can use:

```go
    serviceName := config.Get("service", "name").String("Default Service Name")
```

### Supported Data Types

The following table contains the translation between the viper function signatures and the config functions we have defined.

| Viper Function | Config Function | Return Data Type |
| -------------- | --------------- | ---------------- |
| viper.Get(string) | config.Get(...string).Value(interface{}) | interface{} |
| viper.GetBool(string) | config.Get(...string).Bool(bool) | bool | 
| viper.GetFloat64(string) | config.Get(...string).Float64(float64) | float64 |
| viper.GetInt(string) | config.Get(...string).Int(int) | int |
| viper.GetString(string) | config.Get(...string).String(string) | string |
| viper.GetStringMap(string) | config.Get(...string).StringMap(map[string]interface{}) | map[string]interface{} |
| viper.GetStringMapString(string) | config.Get(...string).StringMapString(map[string]string) | map[string]string |
| viper.GetStringSlice(string) | config.Get(...string).StringSlice([]string) | []string |
| viper.GetTime(string) | config.Get(...string).Time(time.Time) | time.Time |
| viper.GetDuration(string) | config.Get(...string).Duration(time.Duration) | time.Duration |

config.Get takes a variadic string parameter that lays out the path of the configuration you need to retrieve. 
The following type method takes a single parameter that is the default value, which will be returned if the 
configuration is not available in the configuration file.

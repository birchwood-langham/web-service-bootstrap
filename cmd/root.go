// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"gitlab.com/birchwoodlangham/go-web-service-application.git/api"
	"gitlab.com/birchwoodlangham/go-web-service-application.git/config"
	"gitlab.com/birchwoodlangham/go-web-service-application.git/service"

	"github.com/mitchellh/go-homedir"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var application service.Application

var rootCmd = &cobra.Command{
	Run: startService,
}

type logFormatter struct{}

// MaxPort returns the maximum port number available to run your service on
func MaxPort() int {
	return int(^uint16(0))
}

func (f *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := fmt.Sprintf("%s - %s: %s\n", entry.Time.Format("2006-01-02 15:04:05.000"), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func startService(cmd *cobra.Command, args []string) {
	if err := application.Init(); err != nil {
		log.Fatalf("Could not initialize the application -- %s", err)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	checkConfiguration(config.ServiceHostKey, config.ServicePortKey)

	serverHost := viper.GetString(config.ServiceHostKey)
	serverPort := viper.GetInt(config.ServicePortKey)

	maxPort := MaxPort()

	if serverPort == 0 || serverPort > maxPort {
		log.Fatalf("Server port is out of range, port must be between %d and %d", 1, maxPort)
	}

	log.Infof("Starting service on %s:%d", serverHost, serverPort)

	serverMsgChannel := make(chan api.ServerMessage, viper.GetInt(config.ServiceCommandBufferKey))

	go startServer(serverMsgChannel, serverHost, serverPort, application.InitializeRoutes)

	run := true

	for run {
		select {
		case incomingSignal := <-signalChannel:
			log.Infof("Caught signal %v: terminating\n", incomingSignal)

			if err := application.Cleanup(); err != nil {
				log.Errorf("Could not execute cleanup - %s", err)
				continue
			}

			run = false
		case incomingServerMessage := <-serverMsgChannel:
			switch incomingServerMessage {
			case api.Stop:
				log.Info("Stop request from API server has been received, stopping service")
				if err := application.Cleanup(); err != nil {
					log.Errorf("Could not execute cleanup - %s", err)
					continue
				}
				run = false
			default:
				log.Infof("Received an unrecognised command from the API server: %d, ignoring", incomingServerMessage)
			}
		}
	}
}

func startServer(messageChannel chan api.ServerMessage, host string, port int, initializeRoutes func(*api.Server)) {
	server := api.New(host, port, messageChannel)
	server.Initialize(initializeRoutes)
	server.Run()
}

func checkConfiguration(configs ...string) {
	for _, c := range configs {
		if !viper.IsSet(c) {
			log.Fatalf("could not find configuration for: %s, cannot continue", c)
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(app service.Application) {
	application = app

	rootCmd.Use = app.Properties().Usage
	rootCmd.Short = app.Properties().ShortDescription
	rootCmd.Long = app.Properties().LongDescription

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default will search for $PWD/application.yaml then $HOME/application.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".trend-risk" (without extension).
		viper.AddConfigPath(currentDir)
		viper.AddConfigPath(home)
		viper.SetConfigName("application")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		setupLogger()

		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func setupLogger() {
	fileHook := lfshook.NewHook(viper.GetString(config.LogFilePathKey), new(logFormatter))
	log.SetFormatter(new(logFormatter))
	log.SetOutput(os.Stderr)
	log.AddHook(fileHook)

	switch strings.ToUpper(viper.GetString(config.LogLevelKey)) {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

// GetRootCommand returns the service RootCommand so that you can extend it and add your own commands
func GetRootCommand() *cobra.Command {
	return rootCmd
}

// AddCommand adds additional commands to the Root Command
func AddCommand(commands ...*cobra.Command) {
	for _, c := range commands {
		rootCmd.AddCommand(c)
	}
}

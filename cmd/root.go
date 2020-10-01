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
	"syscall"

	"go.uber.org/zap"

	"github.com/birchwood-langham/web-service-bootstrap/api"
	"github.com/birchwood-langham/web-service-bootstrap/config"
	"github.com/birchwood-langham/web-service-bootstrap/logger"
	"github.com/birchwood-langham/web-service-bootstrap/service"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var application service.Application

var rootCmd = &cobra.Command{
	Run: startService,
}

// MaxPort returns the maximum port number available to run your service on
func MaxPort() int {
	return int(^uint16(0))
}

func startService(cmd *cobra.Command, args []string) {
	if err := application.Init(); err != nil {
		zap.S().Fatalf("Could not initialize the application -- %s", err)
	}

	signalChannel := make(chan os.Signal, 100)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	checkConfiguration(config.ServiceHostKey, config.ServicePortKey)

	serverHost := "localhost"

	if viper.IsSet(config.ServiceHostKey) {
		serverHost = viper.GetString(config.ServiceHostKey)
	}

	serverPort := 9900

	if viper.IsSet(config.ServicePortKey) {
		serverPort = viper.GetInt(config.ServicePortKey)
	}

	maxPort := MaxPort()

	if serverPort == 0 || serverPort > maxPort {
		zap.S().Fatalf("Server port is out of range, port must be between %d and %d", 1, maxPort)
	}

	zap.S().Infof("Starting service on %s:%d", serverHost, serverPort)

	serverMsgChannel := make(chan struct{}, viper.GetInt(config.ServiceCommandBufferKey))

	go startServer(serverMsgChannel, serverHost, serverPort, application.InitializeRoutes)

runLoop:
	for {
		select {
		case incomingSignal := <-signalChannel:
			zap.S().Infof("Caught signal %v: terminating\n", incomingSignal)

			serverMsgChannel <- struct{}{}
		case <-serverMsgChannel:
			zap.S().Info("Stop request from API server has been received, stopping service")
			if err := application.Cleanup(); err != nil {
				zap.S().Errorf("Could not execute cleanup - %s", err)
			}

			break runLoop
		}
	}
}

func startServer(messageChannel chan struct{}, host string, port int, initializeRoutes func(*api.Server)) {
	server := api.New(host, port, messageChannel)

	server.Initialize(initializeRoutes)
	server.Run()
}

func checkConfiguration(configs ...string) {
	for _, c := range configs {
		if !viper.IsSet(c) {
			zap.S().Warnf("could not find configuration for: %s, using default values", c)
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
	//initLogger()

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

		// Search config in home directory with name "application" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath("./conf")
		viper.AddConfigPath(home)
		viper.SetConfigName("application")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		zap.S().Errorf("Could not read in viper config: %v", err)
		return
	}

	setupLogger()

	zap.S().Infof("Using config file: %s", viper.ConfigFileUsed())
}

func setupLogger() {
	l := logger.New(logger.ApplicationLogLevel(), logger.DefaultLumberjackLogger())
	zap.ReplaceGlobals(l)
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

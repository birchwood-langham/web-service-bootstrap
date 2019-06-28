package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	// VersionKey is the application.yaml key for retrieving application version number
	VersionKey = "version"
	// ServiceNameKey is the application.yaml key for retrieving the service name
	ServiceNameKey = "service.name"
	// ServiceHostKey is the application.yaml key for retrieving the service host, i.e. the host exposed by the service
	ServiceHostKey = "service.host"
	// ServicePortKey is the application.yaml key for retrieving the port to run the service on
	ServicePortKey = "service.port"
	// ServiceCommandBufferKey is the application.yaml key for retrieving the size of the command buffer for sending control messages to the service thread
	ServiceCommandBufferKey = "service.api-command-buffer"
	// ServiceWriteTimeoutKey is the application.yaml key for retrieving the write timeout for the server
	ServiceWriteTimeoutKey = "service.write-timeout-seconds"
	// ServiceReadTimeoutKey is the application.yaml key for retrieving the read timeout for the server
	ServiceReadTimeoutKey = "service.read-timeout-seconds"
	// ServiceIdleTimeoutKey is the application.yaml key for retrieving the idle timeout for the server
	ServiceIdleTimeoutKey = "service.idle-timeout-seconds"
	// LogFilePathKey is the application.yaml key for retrieving the path for the log file generated by the service
	LogFilePathKey = "log-file-path"
	// LogLevelKey is the application.yaml key for retrieving the logging level
	LogLevelKey = "log-level"
)

type Config struct {
	path []string
}

func Get(path ...string) *Config {
	return &Config{path: path}
}

func (c *Config) String(d string) string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetString(k)
	}

	return d
}

func (c *Config) Int(d int) int {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetInt(k)
	}

	return d
}

func (c *Config) Get(d interface{}) interface{} {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.Get(k)
	}

	return d
}

func (c *Config) GetBool(d bool) bool {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetBool(k)
	}

	return d
}

func (c *Config) GetFloat64(d float64) float64 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetFloat64(k)
	}

	return d
}

func (c *Config) GetStringMap(d map[string]interface{}) map[string]interface{} {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringMap(k)
	}

	return d
}

func (c *Config) GetStringMapString(d map[string]string) map[string]string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringMapString(k)
	}

	return d
}

func (c *Config) GetStringSlice(d []string) []string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringSlice(k)
	}

	return d
}

func (c *Config) GetTime(d time.Time) time.Time {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetTime(k)
	}

	return d
}

func (c *Config) GetDuration(d time.Duration) time.Duration {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetDuration(k)
	}

	return d
}


func mkString(sep string, input ...string) string {
	b := strings.Builder{}

	addSep := false

	for _, i := range input {
		if addSep {
			b.WriteString(sep)
		}

		b.WriteString(i)

		addSep = true
	}

	return b.String()
}
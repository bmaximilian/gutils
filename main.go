package main

import (
	"fmt"
	"github.com/bmaximilian/gutils/cmd"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Version of the cli tool
var Version = "0.0.1"

// Build (date) of the cli tool
var Build = "1"

// Set the default environment variables
// These variables can be overridden by the config file
func setDefaultEnvironment() {
	defaultAppEnv := os.Getenv("APP_ENV")
	if defaultAppEnv == "" {
		defaultAppEnv = "production"
	}

	defaultConfigPathEnv := os.Getenv("ROOM_CALC_CONFIG_DIRECTORY")
	if defaultConfigPathEnv == "" {
		defaultConfigPathEnv = "."
	}

	viper.SetDefault("APP_ENV", defaultAppEnv)
	viper.SetDefault("ROOM_CALC_CONFIG_DIRECTORY", defaultConfigPathEnv)
	viper.SetDefault("VERSION", Version)
	viper.SetDefault("BUILD", Build)

	viper.SetDefault("log.file", "./logs/gutils.log")
	viper.SetDefault("log.level", "INFO")
	cmd.SetDefaults()
}

// Loads the config file
func initConfig() {
	configSuffix := ""
	switch viper.Get("APP_ENV") {
	case "development":
		configSuffix = ".development"
		break
	default:
		break
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(viper.Get("ROOM_CALC_CONFIG_DIRECTORY").(string))
	viper.SetConfigName("gutils.config" + configSuffix)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func main() {
	setDefaultEnvironment()
	initConfig()
	logger.SetLogPath(viper.Get("log.file").(string))

	logLevel := strings.ToUpper(viper.Get("log.level").(string))
	if logger.GetLogLevelForName(logLevel) == nil {
		logLevel = "INFO"
	}

	l := logger.GetLogger()
	l.SetLogLevel(logLevel)
	l.SetForceCli(true)
	// From now on we can use the logger to log our messages

	cmd.Execute()
}

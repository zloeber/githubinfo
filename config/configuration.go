package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/zloeber/githubinfo/log"
)

var defaultConfig *viper.Viper
var cfgFile string

// Configuration for running this application
type Configuration struct {
	Project ProjectConfiguration
}

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

// // ConfigFile returns the current configuration file
// func ConfigFile() string {
// 	conffile := os.LookupEnv()().ExpandEnv("$HOME/.config"),
// 		"Path to config file")
// 	ConfigFile
// }

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)

	var configuration Configuration

	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("config")
		v.AddConfigPath(".")
	}
	v.AutomaticEnv()

	// global defaults
	v.SetDefault("verbose", false)
	v.SetDefault("loglevel", "info")

	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	} else {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := v.Unmarshal(&configuration); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	return v
}

// init this module
func init() {
	defaultConfig = readViperConfig("GITHUBINFO")
}

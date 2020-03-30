package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/zloeber/githubinfo/pkg/version"
)

var (
	defaultConfig *viper.Viper
	cfgFile       string
)

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

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, ", err)
		os.Exit(1)
	}

	if err := v.Unmarshal(&configuration); err != nil {
		fmt.Println("Error unmarshalling config file, ", err)
		os.Exit(1)
	}

	return v
}

// init this module
func init() {
	defaultConfig = readViperConfig(strings.ToUpper(version.AppName))
}

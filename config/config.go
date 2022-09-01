package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

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

var defaultConfig *viper.Viper
var provider *viper.Viper
var once sync.Once

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider() Provider {

	once.Do(func() {
		provider = readViperConfig()
	})
	return provider
}

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	// 环境变量设置支持
	v := viper.New()
	v.SetEnvPrefix("ipumpkin")

	// 文件设置支持
	v.AddConfigPath(".")
	v.AddConfigPath("/home/projects/fyh/ipumpkin")
	v.SetConfigName("settings")
	v.SetConfigType("toml")

	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	// 本地开发配置
	runLevel := v.GetString("run_level")
	if runLevel == "development" {
		v.SetConfigName("settings.local")
		v.SetConfigType("toml")
		v.AddConfigPath(".")
		v.MergeInConfig()
	}

	return v
}
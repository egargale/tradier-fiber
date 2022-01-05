package util

import (
	"fmt"
	// "log"
	"time"

	"github.com/spf13/viper"
)

var MyConfig Config

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	TradierKey            string        `mapstructure:"TRADIER_KEY"`
	TradierAccount        string        `mapstructure:"TRADIER_ACCOUNT"`
	TradierSandboxKey     string        `mapstructure:"TRADIER_SANDBOX_KEY"`
	TradierSandboxAccount string        `mapstructure:"TRADIER_SANDBOX_ACCOUNT"`
	RedisHost             string        `mapstructure:"REDIS_HOST"`
	RedisPort             string        `mapstructure:"REDIS_PORT"`
	DBDriver              string        `mapstructure:"DB_DRIVER"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) error {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&MyConfig)
}

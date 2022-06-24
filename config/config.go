package config

import (
	"errors"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
)

// Config App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

// ServerConfig Server config struct
type ServerConfig struct {
	Port    string
	Swagger string
}

// PostgresConfig Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
}

// LoadConfig Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		klog.Errorf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

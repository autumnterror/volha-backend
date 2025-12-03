package config

import (
	"log"
	"os"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/spf13/viper"
)

type Config struct {
	ConnStr string `mapstructure:"conn_str"`
	Port    int    `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
}

// MustSetup return config and panic if error
func MustSetup() *Config {
	cfg, err := setup()
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

// setup create config structure
func setup() (*Config, error) {
	const op = "config.setup"

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./local-config/blocknote.yaml"
	}

	viper.SetConfigFile(configPath)
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, format.Error(op, err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, format.Error(op, err)
	}

	if cfg.Mode == "DEV" {
		log.Println(format.Struct(cfg))
	}

	return &cfg, nil
}

package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/joho/godotenv"
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

	type _сonfig struct {
		DataSource   string `mapstructure:"data_source"`
		PortPostgres int    `mapstructure:"port_postgres"`
		Port         int    `mapstructure:"port"`
		Mode         string `mapstructure:"mode"`
	}
	if err := godotenv.Load(); err != nil {
		return nil, format.Error(op, err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./local-config/blocknote.yaml"
	}

	viper.SetConfigFile(configPath)
	var cfg _сonfig
	if err := viper.ReadInConfig(); err != nil {
		return nil, format.Error(op, err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, format.Error(op, err)
	}
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")

	if user == "" || pw == "" || db == "" {
		return nil, format.Error(op, errors.New("missing environment variables"))
	}

	if cfg.Mode == "DEV" {
		log.Println(format.Struct(cfg), fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			user, pw, cfg.DataSource, cfg.PortPostgres, db))
	}

	return &Config{
		ConnStr: fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			user, pw, cfg.DataSource, cfg.PortPostgres, db),
		Port: cfg.Port,
		Mode: cfg.Mode,
	}, nil
}

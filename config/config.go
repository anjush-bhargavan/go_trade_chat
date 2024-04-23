package config

import "github.com/spf13/viper"

// Config represents the environment variables.
type Config struct {
	SECRETKEY string `mapstructure:"JWTSECRET"`
	GrpcPort      string `mapstructure:"GRPCPORT"`
	Host         string `mapstructure:"HOST"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Database     string `mapstructure:"DBNAME"`
	Port         string `mapstructure:"PORT"`
	Sslmode      string `mapstructure:"SSL"`
}

// LoadConfig will load the environment variable to access.
func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}

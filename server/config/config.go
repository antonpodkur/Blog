package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
	Jwt    JwtConfig
}

type ServerConfig struct {
	Port string
}

type MongoConfig struct {
	MongoURI string
    DbName string
}

type JwtConfig struct {
	AccessTokenPrivateKey string
	AccessTokenPublicKey  string
	AccessTokenExpiresIn  time.Duration
	AccessTokenMaxAge     int

	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	RefreshTokenExpiresIn  time.Duration
	RefreshTokenMaxAge     int
}

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

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

func GetConfigPath(configPath string) string {
   if configPath == "docker" {
       return "./config/config-docker"
   }
   return "./config/config-local"
}

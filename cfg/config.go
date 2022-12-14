package cfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	SERVER   Server
	DATABASE Database
}

type Server struct {
	PORT      string
	VERSION   string
	ENV       string
	EMAIL     string
	EMAIL_PWD string
	JWT       Jwt
}

type Jwt struct {
	REFRESH_COOKIE_NAME string
	ACCESS_COOKIE_NAME  string
	REFRESH_EXPIRES_AT  int32
	ACCESS_EXPIRES_AT   int32
	SECRET_KEY_ACCESS   string
	SECRET_KEY_REFRESH  string
}

type Database struct {
	PORT   string
	HOST   string
	NAME   string
	PSWD   string
	USER   string
	DRIVER string
}

func New(name string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("cfg")
	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("could not read in config: %v", err)
	}

	return v, nil
}

func Parse(v *viper.Viper) (*Config, error) {
	var cfg *Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("could not unmarsh into config: %v", err)
	}

	return cfg, nil
}
